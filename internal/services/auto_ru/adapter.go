package auto_ru

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/tonatos/goal-tracker/internal/services"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

type AutoRUApiURLs struct {
	Catalog     string
	ApiCatalog  string
	ApiAdsCount string
}

type AutoRU struct {
	Amount       float32
	Urls         AutoRUApiURLs
	Api          services.Api
	FilterParams AutoRUCountAdsRequest
}

func (ar *AutoRU) Auth() error {
	return nil
}

func (ar *AutoRU) CountAds() (int, error) {
	// Redis key format: `ads:amount__geo_id__radius:count`
	key := utils.CreateRedisKey(
		"ads",
		fmt.Sprintf(
			"%.0f__%d__%d",
			ar.FilterParams.PriceTo,
			ar.FilterParams.GeoID,
			ar.FilterParams.GeoRadius,
		),
		"count",
	)
	redis_val, err := database.Redis.Get(key).Result()
	if err == nil {
		res, _ := strconv.ParseInt(redis_val, 10, 32)
		log.Printf("[INFO] Use cached ads counter (%s): %d", database.Redis.TTL(key), res)
		return int(res), nil
	}

	jsonBody, _ := json.Marshal(ar.FilterParams)
	data, err := ar.Api.Post(ar.Api.BuildURL(ar.Urls.ApiAdsCount), jsonBody)

	if err != nil {
		return 0, err
	}

	var responseObject []AutoRUCountAdsResponse
	json.Unmarshal(data, &responseObject)
	if len(responseObject) == 0 {
		return 0, nil
	}

	var count int
	for _, countByRegion := range responseObject {
		if countByRegion.Radius <= ar.FilterParams.GeoRadius {
			count = countByRegion.Count
		}
	}

	err = database.Redis.Set(key, count, time.Second*60*60*6).Err()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ar *AutoRU) GetCatalogLink() (string, error) {
	baseUrl := ar.Api.BuildURL(ar.Urls.Catalog)
	params := url.Values{
		"in_stock": []string{"ANY_STOCK"},
	}

	values := reflect.ValueOf(ar.FilterParams)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		name := types.Field(i).Tag.Get("json")
		value := values.Field(i).Interface()
		switch v := value.(type) {
		case string:
			params.Add(name, value.(string))
		case int:
			params.Add(name, strconv.Itoa(value.(int)))
		case float32:
			params.Add(name, fmt.Sprintf("%.0f", value.(float32)))
		case bool:
			params.Add(name, strconv.FormatBool(value.(bool)))
		case []map[string]string:
			for _, v := range value.([]map[string]string) {
				params.Add(
					name,
					fmt.Sprintf("mark=%s,model=%s", v["mark"], v["model"]),
				)
			}
		default:
			log.Printf("[WARN] Value %s %+v can't converted\n", name, v)
		}

	}
	baseUrl.RawQuery = params.Encode()
	return baseUrl.String(), nil
}

func AutoruInit() {
	cookie := os.Getenv("APP_AUTORU_COOKIE")

	AutoruObject = &AutoRU{
		Urls: AutoRUApiURLs{
			ApiAdsCount: "/-/ajax/desktop/getListingLocatorCountersTotalCount/",
			ApiCatalog:  "/-/ajax/desktop/listing/",
			Catalog:     "/ekaterinburg/cars/used/",
		},
		FilterParams: AutoRUCountAdsRequest{
			WithDiscount: false,
			Section:      "used",
			Category:     "cars",
			GeoRadius:    1000,
			GeoID:        54,
			CatalogFilter: []map[string]string{
				{"mark": "CHEVROLET", "model": "CAMARO"},
				{"mark": "FORD", "model": "MUSTANG"},
				{"mark": "DODGE", "model": "CHALLENGER"},
			},
		},
		Api: &AutoRUApi{
			BaseURL: "https://auto.ru",
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "*/*",
				"User-Agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Safari/605.1.15",
				"Cookie":       cookie,
			},
		},
	}
}

func AutoruSetupObject(goal_amount, accumulatedAmount float32) {
	AutoruObject.Amount = goal_amount
	AutoruObject.FilterParams.PriceTo = accumulatedAmount
}

var AutoruObject *AutoRU
