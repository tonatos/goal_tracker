package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

func createMockImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}
	return img
}

func (suite *HandlersTestSuite) TestUploadImage() {
	imageName := "someimg.png"

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", imageName)
	if err != nil {
		suite.Suite.Error(err)
	}

	err = png.Encode(part, createMockImage())
	if err != nil {
		suite.Suite.Error(err)
	}
	writer.Close()

	tests := []struct {
		name                 string
		method               string
		url                  string
		requestData          io.Reader
		expectedError        bool
		expectedResponseData map[string]interface{}
		expectedResponseCode int
	}{
		{
			name:                 "Test Upload: basic",
			method:               "POST",
			url:                  "upload/",
			requestData:          body,
			expectedError:        false,
			expectedResponseCode: 200,
			expectedResponseData: map[string]interface{}{
				"code":    float64(200),
				"message": "success",
			},
		},
		{
			name:                 "Test Upload: empty",
			method:               "POST",
			url:                  "upload/",
			requestData:          nil,
			expectedResponseCode: 400,
			expectedResponseData: map[string]interface{}{
				"code":    float64(400),
				"message": "error",
			},
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, fmt.Sprintf("%s%s", suite.Api, tt.url), tt.requestData)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			resp, err := suite.App.Test(req, -1)

			// Verify, that no error occurred, that is not expected
			assert.Equalf(suite.T(), tt.expectedError, err != nil, tt.name)

			// Verify, if the status code is as expected.
			assert.Equalf(suite.T(), tt.expectedResponseCode, resp.StatusCode, tt.name)

			if tt.expectedResponseCode == 200 {
				var responseJson map[string]interface{}
				responseBody, _ := io.ReadAll(resp.Body)
				if err := json.Unmarshal(responseBody, &responseJson); err != nil {
					t.Errorf("%s: %s", tt.name, err)
				}

				imagepath := fmt.Sprintf(
					"%s/%s",
					utils.GetBaseDir(),
					responseJson["data"].(map[string]interface{})["image"].(string),
				)

				if _, err := os.Stat(imagepath); err != nil {
					assert.NoError(suite.T(), err)
				}

				if err := os.Remove(imagepath); err != nil {
					assert.NoError(suite.T(), err)
				}
			}

			defer resp.Body.Close()
			defer req.Body.Close()
		})
	}
}
