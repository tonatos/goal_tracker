package database

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Slug       string    `gorm:"unique" json:"slug"`
	GoalAmount float32   `json:"goal_amount" sql:"type:decimal(10,2);"`
	TargetDate time.Time `json:"target_date"`
}

type Contribution struct {
	gorm.Model `json:"-"`
	ID         uint      `gorm:"primaryKey"`
	GoalID     uint      `json:"goal_id"`
	Goal       Goal      `json:"goal"`
	Amount     float32   `json:"amount" sql:"type:decimal(10,2);"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Contribution) TableName() string {
	return "contributions"
}
