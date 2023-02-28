package request

import (
	"time"
)

type RequestCreateGoal struct {
	Name       string    `json:"name,omitempty"`
	GoalAmount float32   `json:"goal_amount,omitempty" sql:"type:decimal(10,2);"`
	TargetDate time.Time `json:"target_date,omitempty"`
}

type RequestUpdateGoal struct {
	*RequestCreateGoal
}

type RequestCreateContribution struct {
	Amount float32 `json:"amount,omitempty" sql:"type:decimal(10,2);"`
}

type RequestUpdateContribution struct {
	Amount float32 `json:"amount,omitempty" sql:"type:decimal(10,2);"`
}
