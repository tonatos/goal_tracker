package utils

import "fmt"

func CreateRedisKey(object, id, property string) string {
	return fmt.Sprintf("%s:%s:%s", object, id, property)
}
