package fixtures

import (
	"github.com/tonatos/goal-tracker/internal/models/table"
)

type ContributionFixture struct {
	Items []table.Contribution
}

func ContributionFixturesFabric(goal *table.Goal) *ContributionFixture {
	contributionFixture := &ContributionFixture{
		Items: []table.Contribution{
			{
				GoalID: goal.ID,
				Amount: 100.0,
			},
			{
				GoalID: goal.ID,
				Amount: 50.0,
			},
		},
	}
	return contributionFixture
}
