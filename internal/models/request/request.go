package request

import (
	"time"
)

type RequestCreateGoal struct {
	Name       string    `json:"name,omitempty" validate:"required"`
	GoalAmount float32   `json:"goal_amount,omitempty" validate:"required"`
	TargetDate time.Time `json:"target_date,omitempty" validate:"required"`
	Image      string    `json:"image,omitempty"`
}

type RequestUpdateGoal struct {
	Name       string    `json:"name,omitempty"`
	GoalAmount float32   `json:"goal_amount,omitempty"`
	TargetDate time.Time `json:"target_date,omitempty"`
	Image      string    `json:"image,omitempty"`
}

type RequestCreateContribution struct {
	Amount float32 `json:"amount,omitempty" validate:"required"`
}

type RequestUpdateContribution struct {
	Amount float32 `json:"amount,omitempty" validate:"required"`
}
