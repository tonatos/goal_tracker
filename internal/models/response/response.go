package response

import (
	"time"

	"github.com/tonatos/goal-tracker/internal/models/table"
)

type ResponesGoal struct {
	*table.Goal
	AccumulatedAmount float32 `json:"accumulated_amount"`
	CatalogUrl        string  `json:"catalog_url"`
	AdsByAmount       int     `json:"ads_by_amount"`
	DaysUntilBang     int     `json:"days_until_bang"`
}

type ResponesContribution struct {
	Id        uint      `json:"id"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
