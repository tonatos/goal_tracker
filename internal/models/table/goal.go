package table

import (
	"time"
)

type Goal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Slug       string    `gorm:"unique" json:"slug"`
	GoalAmount float32   `gorm:"not null" json:"goal_amount" sql:"type:decimal(10,2);"`
	TargetDate time.Time `gorm:"not null" json:"target_date"`
	Image      string    `json:"image"`
}
