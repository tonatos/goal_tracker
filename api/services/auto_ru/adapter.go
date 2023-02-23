package auto_ru

import (
	"encoding/json"
	"fmt"
	"goal-tracker/api/services"
	"net/url"
	"os"
	"reflect"
	"strconv"
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
		switch value.(type) {
		case string:
			params.Add(name, value.(string))
		case int:
			params.Add(name, strconv.Itoa(value.(int)))
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
			break
		}

	}
	baseUrl.RawQuery = params.Encode()
	return baseUrl.String(), nil
}

func AutoruInit(goal_amount float32) *AutoRU {
	cookie := os.Getenv("APP_AUTORU_COOKIE")

	return &AutoRU{
		Amount: goal_amount,
		Urls: AutoRUApiURLs{
			ApiAdsCount: "/-/ajax/desktop/getListingLocatorCountersTotalCount/",
			ApiCatalog:  "/-/ajax/desktop/listing/",
			Catalog:     "/ekaterinburg/cars/used/",
		},
		FilterParams: AutoRUCountAdsRequest{
			WithDiscount: false,
			Section:      "used",
			Category:     "cars",
			GeoRadius:    200,
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
