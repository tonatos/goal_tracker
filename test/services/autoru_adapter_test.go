package services

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonatos/goal-tracker/internal/services/auto_ru"
	"github.com/tonatos/goal-tracker/test/mocks"
)

func (suite *ServicesTestSuite) TestAutoRUAdapter() {

	tests := []struct {
		name             string
		amount           float32
		expectedAdsCount int
		expectedLink     string
		filter           auto_ru.AutoRUCountAdsRequest
		postResponse     []byte
		postError        error
		wantError        bool
	}{
		{
			name:             "TestAutoRUAdapter: base",
			amount:           1000.0,
			expectedAdsCount: 10,
			expectedLink:     "https://testurl.com/catalog/?catalog_filter=mark%3DFORD%2Cmodel%3DMUSTANG&category=cars&geo_id=1&geo_radius=1000&in_stock=ANY_STOCK&price_to=100&section=used&with_discount=false",
			filter: auto_ru.AutoRUCountAdsRequest{
				WithDiscount: false,
				Section:      "used",
				Category:     "cars",
				GeoRadius:    1000,
				GeoID:        1,
				PriceTo:      100.0,
				CatalogFilter: []map[string]string{
					{"mark": "FORD", "model": "MUSTANG"},
				},
			},
			postResponse: []byte(`[{"radius": 100, "count": 10}]`),
		},
		{
			name:             "TestAutoRUAdapter: different radius",
			amount:           1000.0,
			expectedAdsCount: 10,
			expectedLink:     "https://testurl.com/catalog/?catalog_filter=mark%3DFORD%2Cmodel%3DMUSTANG&category=cars&geo_id=1&geo_radius=500&in_stock=ANY_STOCK&price_to=500&section=used&with_discount=false",
			filter: auto_ru.AutoRUCountAdsRequest{
				WithDiscount: false,
				Section:      "used",
				Category:     "cars",
				GeoRadius:    500,
				GeoID:        1,
				PriceTo:      500.0,
				CatalogFilter: []map[string]string{
					{"mark": "FORD", "model": "MUSTANG"},
				},
			},
			postResponse: []byte(`[{"radius": 500, "count": 10},{"radius": 1000, "count": 50}]`),
		},
		{
			name:             "TestAutoRUAdapter: money is too small",
			amount:           1000.0,
			expectedAdsCount: 0,
			expectedLink:     "https://testurl.com/catalog/?catalog_filter=mark%3DFORD%2Cmodel%3DMUSTANG&category=cars&geo_id=1&geo_radius=1000&in_stock=ANY_STOCK&price_to=0&section=used&with_discount=false",
			filter: auto_ru.AutoRUCountAdsRequest{
				WithDiscount: false,
				Section:      "used",
				Category:     "cars",
				GeoRadius:    1000,
				GeoID:        1,
				PriceTo:      0.0,
				CatalogFilter: []map[string]string{
					{"mark": "FORD", "model": "MUSTANG"},
				},
			},
			postResponse: []byte(`[]`),
		},
		{
			name:             "TestAutoRUAdapter: error when filter not passed",
			amount:           1000.0,
			expectedAdsCount: 0,
			expectedLink:     "https://testurl.com/catalog/?category=&geo_id=0&geo_radius=0&in_stock=ANY_STOCK&price_to=0&section=&with_discount=false",
			filter:           auto_ru.AutoRUCountAdsRequest{},
			postResponse:     []byte(`[]`),
			postError:        errors.New("request error"),
			wantError:        true,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			api := mocks.MockAutoruApiInit(false)

			body, _ := json.Marshal(&tt.filter)
			api.On("Post", api.BuildURL("/api/ads_count/"), body).Return(tt.postResponse, tt.postError)
			autoru := mocks.MockAutoruInit(api, &tt.filter, tt.amount, tt.filter.PriceTo)

			ads, err := autoru.CountAds()

			// Verify, that no error occurred, that is not expected
			assert.Equalf(suite.T(), tt.wantError, err != nil, tt.name)

			// Verify value
			assert.Equalf(suite.T(), tt.expectedAdsCount, ads, tt.name)

			link, _ := autoru.GetCatalogLink()

			// Verify builded link
			assert.Equalf(suite.T(), tt.expectedLink, link, tt.name)
		})
	}
}
