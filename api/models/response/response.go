package response

import (
	"goal-tracker/api/models/db"
)

type Goal struct {
	*db.Goal
	CatalogUrl  string `json:"catalog_url"`
	AdsByAmount int    `json:"ads_by_amount"`
}
