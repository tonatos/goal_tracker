package utils

import "encoding/json"

func StructToMap(source interface{}) map[string]interface{} {
	var result map[string]interface{}
	inrec, _ := json.Marshal(source)
	json.Unmarshal(inrec, &result)
	return result
}
