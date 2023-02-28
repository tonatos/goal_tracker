package table

import (
	"time"

	"gorm.io/gorm"
)

type Contribution struct {
	gorm.Model `json:"-"`
	ID         uint      `gorm:"primaryKey" json:"id"`
	GoalID     uint      `json:"goal_id"`
	Goal       Goal      `gorm:"constraint:OnDelete:CASCADE" json:"goal"`
	Amount     float32   `json:"amount" sql:"type:decimal(10,2);"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Contribution) TableName() string {
	return "contributions"
}
