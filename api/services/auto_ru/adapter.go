package auto_ru

import (
	"encoding/json"
	"fmt"
	"goal-tracker/api/services"
	"net/url"
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
				"Cookie":       "from=direct; from_lifetime=1677178681917; gdpr=0; yaPlusUserPlusBalance={%22id%22:%2223758630%22%2C%22plusPoints%22:822}; count-visits=7; sso_status=sso.passport.yandex.ru:synchronized_no_beacon; _ym_d=1677178678; _ym_uid=1674634997900228470; counter_ga_all7=2; layout-config={\"screen_height\":982,\"screen_width\":1512,\"win_width\":1512,\"win_height\":438}; cycada=z/wJboN82ZmifbfnMnJZZRT8KkOa6Byr+ShFrwT+oxc=; spravka=dD0xNjc3MTc4NjM2O2k9MTg4LjI0NC4xNzYuMjUwO0Q9M0JENTJFQTgzMjU5QUNCQUE1NzI3RDBDNEE3NDQzNTY2MTA5RTgxQ0NDQTBGNEIxNzFCN0E2MTIzMDFCOTYzMENEQURGMEM3MjY7dT0xNjc3MTc4NjM2MDgxMzM4ODI1O2g9ZGE1YTRkMTQ4NTdkNGZjZGY5ZTI2OTJmYzNhZjM2MDQ=; _ym_visorc=b; _yasc=bzOT7idVUBY0KYTo7ThMu9jX9yfEDX+ybktyqLtoE5Rx+h+R7KKSSit3PEb/yOYpfg==; _ym_isad=2; autoru-visits-count=10; _csrf_token=c5eda521961a0236f76601ef6887da912b309b6d54567fc7; my=YycCAAEA; yandexuid=5296136751638362801; ys=udn.cDpUb25hdG9zRm9uVGllcg%3D%3D%23c_chck.1103605736; yuidlt=1; popups-autoru-plus-shown-count=1; popups-autoru-only-shown-count=1; autoruuid=g63d0e6fd2fh4rh3k3vo3o9gviana1nu.692c7a875548129cef78a7c430452afc; autoru_sid=23758630%7C1675401000671.7776000.MPvWXVLT4E3Ddka_grS--A.2xCEXouv91evoaxlBqoD-hBXbj_5NQwOQSNwdSkfKzo; L=emleSHF7V39lQAFaeH1EQU0BVkN/YWRyZgY+Kj8ORwcIHRkrE0Y=.1673515772.15220.322486.9d04247b1ee14d5400b1647048efbcd6; Session_id=3:1675400984.5.1.1654757435462:L4REvA:6.1.2:1|1130000057383267.0.2|20965903.13143123.2.2:13143123|61:10010785.712835.MnqC-1lDqPVArWzRTfXIkYg78ig; i=fWgh0Nxqe+GABlyCHUR1EgZHZJUHwfntav16MWcI9tRC5LvRGrhRq1YbaVeFnY9YG8PL7o1Pzqh/nyRrnOPjZQpF2bs=; mda2_beacon=1675400984933; yandex_login=TonatosFonTier; safe_deal_promo=3; gids=54; gradius=200; suid=f814d89e43034b9578644e5b3181dee1.84b590f2bfb283ae10b078de42263b34",
			},
		},
	}
}
