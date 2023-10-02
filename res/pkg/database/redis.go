package database

const Redis = `package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"{{ .ModuleName}}/pkg/config"
	"{{ .ModuleName}}/pkg/logger"
)

type RedisDB struct {
	Client *redis.Client
}

var (
	Redis *RedisDB
)

func ConnetRedis() *RedisDB {
	cfg := config.RedisDBCfg()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), // Redis server address
		Password: cfg.Password,                             // Password (if any)
		Username: cfg.User,
		DB:       0, // Redis database index
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		logger.Error("failed to ping the redis server:", err)
		return nil
	}
	fmt.Println("Connected to Redis:", pong)

	Redis = &RedisDB{
		Client: client,
	}
	return Redis
}

// AddMapDataInRedis 
func (r *RedisDB) AddMapDataInRedis(key string, data interface{}) error {
	logger.Info(fmt.Sprintf("storing the map data %s to %v", key, data))

	err := r.Client.HSet(context.Background(), key, data).Err()
	if err != nil {
		logger.Error("failed to store map data:", err)
		return err
	}
	return nil
}

// GetMapDataFromRedis
func (r *RedisDB) GetMapDataFromRedis(key string) (map[string]interface{}, error) {
	logger.Info(fmt.Sprintf("getting the map data %s", key))

	val, err := r.Client.HGetAll(context.Background(), key).Result()
	if err != nil {
		logger.Error("failed to get map data:", err)
		return nil, err
	}
	return val, nil
}


// DeleteKeyByKey deletes the key from redis
func (r *RedisDB) DeleteKey(key string) error {
	logger.Info(fmt.Sprintf("Deleting the key %s", key))

	err := r.Client.Del(r.Client.Context(), key).Err()
	if err != nil {
		logger.Error("failed to delete the key:", err)
		return err
	}
	return nil
}
`
