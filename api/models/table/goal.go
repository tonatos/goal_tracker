package table

import (
	"time"
)

type Goal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Slug       string    `gorm:"unique" json:"slug"`
	GoalAmount float32   `json:"goal_amount" sql:"type:decimal(10,2);"`
	TargetDate time.Time `json:"target_date"`
}
