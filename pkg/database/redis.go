package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

type ConfigRedis struct {
	ServerName string
	Port       string
	Password   string
	DB         int
}

func ConnectRedis(config *ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.ServerName, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

var Redis *redis.Client
