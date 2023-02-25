package utils

import "os"

func GetEnvWithDefault(key string, default_value string) string {
	env, finded := os.LookupEnv(key)
	if !finded {
		return default_value
	}
	return env
}
