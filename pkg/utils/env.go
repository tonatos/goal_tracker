package utils

import "os"

func GetCurrentEnv() string {
	var env string
	switch os.Getenv("ENV") {
	case "prod", "stage":
		env = "prod"
	case "test":
		env = "test"
	default:
		env = "dev"
	}
	return env
}

func GetEnvWithDefault(key string, default_value string) string {
	env, finded := os.LookupEnv(key)
	if !finded {
		return default_value
	}
	return env
}
