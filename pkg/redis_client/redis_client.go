package redis_client

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

// RedisClient represents a client uses REDIS.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient contruct
func NewRedisClient(host string, db int, password string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisClient{client: client}
}

// Client returns redis.Client.
func (h *RedisClient) Client() *redis.Client {
	return h.client
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (h *RedisClient) Get(key string) string {
	response, err := h.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	}

	return response
}

// Set Redis `SET key value [expiration]` command.
func (h *RedisClient) Set(key string, value interface{}, expiration time.Duration) {
	err := h.client.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
}

// HGetAll Redis HGETALL command is used to get all fields and values of the hash stored at key.
func (h *RedisClient) HGetAll(key string) map[string]string {
	response, err := h.client.HGetAll(key).Result()
	if err != nil {
		panic(err)
	}

	return response
}

// HSet
func (h *RedisClient) HSet(key string, values ...interface{}) {
	err := h.client.HSet(key, values).Err()
	if err != nil {
		panic(err)
	}
}
