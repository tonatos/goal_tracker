package handlers

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type HandlersTestSuite struct {
	suite.Suite
	App        *fiber.App
	Api        string
	DB         *gorm.DB
	TestDBName string
}

func compareMap(t *testing.T, name string, response, expected map[string]interface{}) {
	for k, v := range expected {
		switch v.(type) {
		case map[string]interface{}:
			compareMap(t, name, response[k].(map[string]interface{}), expected[k].(map[string]interface{}))
		default:
			assert.Equalf(t, expected[k], response[k], name)
		}
	}
}

func CompareRRFieldEntry(t *testing.T, name string, body io.ReadCloser, expected map[string]interface{}) {
	var responseJson map[string]interface{}
	responseBody, _ := io.ReadAll(body)
	if err := json.Unmarshal(responseBody, &responseJson); err != nil {
		t.Errorf("%s: %s", name, err)
	}
	compareMap(t, name, responseJson, expected)
}
