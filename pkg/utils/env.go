package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetBaseDir() string {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	abs_fname, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(filepath.Dir(filepath.Dir(abs_fname)))
}

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
