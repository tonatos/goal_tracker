package response

import (
	"goal-tracker/api/models/table"
	"time"
)

type ResponesGoal struct {
	*table.Goal
	CatalogUrl  string `json:"catalog_url"`
	AdsByAmount int    `json:"ads_by_amount"`
}

type ResponesContribution struct {
	Id        uint      `json:"id"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
