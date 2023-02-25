package auto_ru

type AutoRUCountAdsRequest struct {
	WithDiscount  bool                `json:"with_discount"`
	CatalogFilter []map[string]string `json:"catalog_filter"`
	Section       string              `json:"section"`
	Category      string              `json:"category"`
	GeoRadius     int                 `json:"geo_radius"`
	GeoID         int                 `json:"geo_id"`
}

type AutoRUCountAdsResponse struct {
	Radius int `json:"radius"`
	Count  int `json:"count"`
}
