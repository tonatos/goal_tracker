package fixtures

import (
	"time"

	"github.com/tonatos/goal-tracker/internal/models/table"
)

type GoalFixture struct {
	Items []table.Goal
}

func GoalFixturesFabric() *GoalFixture {
	goalFixture := &GoalFixture{
		Items: []table.Goal{
			{
				Name:       "Ford Mustang",
				GoalAmount: 1000.0,
				TargetDate: time.Now().AddDate(0, 1, 0),
			},
		},
	}
	return goalFixture
}
