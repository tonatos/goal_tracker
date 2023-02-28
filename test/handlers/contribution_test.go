package handlers

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *HandlersTestSuite) TestContribution() {
	tests := []struct {
		name                 string
		method               string
		url                  string
		requestData          map[string]interface{}
		expectedError        bool
		expectedResponseData map[string]interface{}
		expectedResponseCode int
	}{
		// Create Contribution
		{
			name:                 "TestCreateContribution: base",
			method:               "POST",
			url:                  "goal/1/contribution/",
			requestData:          map[string]interface{}{"amount": 1.0},
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},
		{
			name:                 "TestCreateContribution: empty request",
			method:               "POST",
			url:                  "goal/1/contribution/",
			requestData:          nil,
			expectedResponseCode: 400,
			expectedResponseData: map[string]interface{}{
				"code":    float64(400),
				"message": "error",
			},
		},

		// Get Contributions
		{
			name:                 "TestGetContributions: base",
			method:               "GET",
			url:                  "goal/1/contribution/",
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},

		// Update Contribution
		{
			name:                 "TestUpdateContribution: base",
			method:               "PUT",
			url:                  "goal/1/contribution/1/",
			requestData:          map[string]interface{}{"amount": float64(10.0)},
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
				"data": map[string]interface{}{
					"amount": float64(10.0),
				},
			},
		},

		// Delete Contribution
		{
			name:                 "TestDeleteContribution: base",
			method:               "DELETE",
			url:                  "goal/1/contribution/3/",
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
