package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (suite *HandlersTestSuite) TestGetContribution() {}

func (suite *HandlersTestSuite) TestGetContributions() {}

func (suite *HandlersTestSuite) TestCreateContributions() {
	goal_data, err := json.Marshal(map[string]interface{}{
		"name":        "test 12edwqqdd",
		"goal_amount": 1.0,
		"target_date": "2024-04-20T00:00:00.511Z",
	})
	assert.NoError(suite.T(), err)

	contribution_data, err := json.Marshal(map[string]float32{"amount": 1.0})
	assert.NoError(suite.T(), err)

	req := httptest.NewRequest("POST", "/api/v1/goal/", bytes.NewBuffer(goal_data))
	req.Header.Set("Content-type", "application/json")
	resp, _ := suite.App.Test(req)
	assert.Equal(suite.T(), "200 OK", resp.Status)
	fmt.Println(resp)

	req = httptest.NewRequest("POST", "/api/v1/goal/1/contribution/", bytes.NewBuffer(contribution_data))
	req.Header.Set("Content-type", "application/json")

	resp, _ = suite.App.Test(req)
	assert.Equal(suite.T(), "200 OK", resp.Status)

	defer resp.Body.Close()
	defer req.Body.Close()
}

func (suite *HandlersTestSuite) TestUpdateContributions() {}

func (suite *HandlersTestSuite) TestDeleteContributions() {}
