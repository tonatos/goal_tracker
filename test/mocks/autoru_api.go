package mocks

import (
	"fmt"
	"net/url"

	"github.com/stretchr/testify/mock"
	"github.com/tonatos/goal-tracker/internal/services/auto_ru"
)

type MockedAutoRuApi struct {
	mock.Mock

	BaseURL string
	Headers map[string]string
}

func (a *MockedAutoRuApi) BuildURL(path string) *url.URL {
	baseUrl, _ := url.Parse(fmt.Sprintf("%s%s", a.BaseURL, path))
	return baseUrl
}

func (a *MockedAutoRuApi) Post(path *url.URL, body []byte) ([]byte, error) {
	args := a.Called(path, body)
	return args.Get(0).([]byte), args.Error(1)
}

func (a *MockedAutoRuApi) Get(path *url.URL) ([]byte, error) {
	args := a.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func MockAutoruApiInit(initDummy bool) *MockedAutoRuApi {
	api := &MockedAutoRuApi{
		BaseURL: "https://testurl.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Cookie":       "_csrf_token=hui",
		},
	}
	if initDummy {
		api.On("Get", mock.Anything).Return([]byte(`[{"result": 0}]`), nil)
		api.On("Post", mock.Anything, mock.Anything).Return([]byte(`[{"radius": 100, "count": 10}]`), nil)
	}
	return api
}

func MockAutoruInit(api *MockedAutoRuApi, filterParams *auto_ru.AutoRUCountAdsRequest, goal_amount, accumulatedAmount float32) *auto_ru.AutoRU {
	if filterParams == nil {
		filterParams = &auto_ru.AutoRUCountAdsRequest{
			WithDiscount: false,
			Section:      "used",
			Category:     "cars",
			GeoRadius:    1000,
			GeoID:        1,
			PriceTo:      accumulatedAmount,
			CatalogFilter: []map[string]string{
				{"mark": "MARK", "model": "MODEL"},
			},
		}
	}

	autoruObject := &auto_ru.AutoRU{
		Amount: goal_amount,
		Urls: auto_ru.AutoRUApiURLs{
			ApiAdsCount: "/api/ads_count/",
			ApiCatalog:  "/api/catalog/",
			Catalog:     "/catalog/",
		},
		FilterParams: *filterParams,
		Api:          api,
	}
	return autoruObject
}
