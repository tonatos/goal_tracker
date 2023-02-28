package handlers

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *HandlersTestSuite) TestGoal() {
	tests := []struct {
		name                 string
		method               string
		url                  string
		requestData          map[string]interface{}
		expectedError        bool
		expectedResponseData map[string]interface{}
		expectedResponseCode int
	}{
		// Create Goal
		{
			name:   "TestCreateGoal: base",
			method: "POST",
			url:    "goal/",
			requestData: map[string]interface{}{
				"name":        "test goal #1",
				"goal_amount": 1.0,
				"target_date": "2024-04-20T00:00:00.511Z",
			},
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},
		{
			name:   "TestCreateGoal: dublicate request",
			method: "POST",
			url:    "goal/",
			requestData: map[string]interface{}{
				"name":        "test goal #1",
				"goal_amount": 1.0,
				"target_date": "2024-04-20T00:00:00.511Z",
			},
			expectedResponseCode: 400,
		},
		{
			name:   "TestCreateGoal: without required fields",
			method: "POST",
			url:    "goal/",
			requestData: map[string]interface{}{
				"name": "test goal #2",
			},
			expectedResponseCode: 400,
			expectedResponseData: map[string]interface{}{
				"code":    float64(400),
				"message": "error",
			},
		},

		// Get Goals
		{
			name:                 "TestGetGoals: base",
			method:               "GET",
			url:                  "goal/",
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},

		// Get Goal by Id
		{
			name:                 "TestGetGoalById: base",
			method:               "GET",
			url:                  "goal/1/",
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},
		{
			name:                 "TestGetGoalById: 404",
			method:               "GET",
			url:                  "goal/100/",
			expectedResponseCode: 404,
			expectedResponseData: map[string]interface{}{
				"code":    float64(404),
				"message": "Can't find goal with this id",
			},
		},

		// Update Goal
		{
			name:   "TestUpdateGoal: base",
			method: "PUT",
			url:    "goal/1/",
			requestData: map[string]interface{}{
				"goal_amount": float64(100.0),
			},
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
				"data": map[string]interface{}{
					"goal_amount": float64(100.0),
				},
			},
		},

		// Delete Goal
		{
			name:                 "TestDeleteGoal: base",
			method:               "DELETE",
			url:                  "goal/2/",
			expectedResponseCode: 200,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			// Prepare request data (for POST, PUT)
			data, err := json.Marshal(tt.requestData)
			assert.NoError(suite.T(), err)

			req := suite.Request(tt.method, tt.url, bytes.NewBuffer(data))
			resp, err := suite.App.Test(req)

			// Verify, that no error occurred, that is not expected
			assert.Equalf(suite.T(), tt.expectedError, err != nil, tt.name)

			// Verify, if the status code is as expected.
			assert.Equalf(suite.T(), tt.expectedResponseCode, resp.StatusCode, tt.name)

			// Verify, that have key fields in response
			CompareRRFieldEntry(suite.T(), tt.name, resp.Body, tt.expectedResponseData)

			defer resp.Body.Close()
			defer req.Body.Close()
		})
	}
}
