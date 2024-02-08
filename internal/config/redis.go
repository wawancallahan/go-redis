package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisConfig struct {
	client *redis.Client
}

func NewRedisConfig() *RedisConfig {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis cannot communicate! ", err)
	}

	return &RedisConfig{
		client: redisClient,
	}
}

func (c *RedisConfig) Set(key string, value string, ttl time.Duration) error {
	err := c.client.Set(ctx, key, value, ttl).Err()

	return err
}

func (c *RedisConfig) Get(key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return value, nil
}

func (c *RedisConfig) Delete(key string) error {
	err := c.client.Del(ctx, key).Err()

	return err
}

func (c *RedisConfig) HSet(key string, value ...interface{}) error {
	err := c.client.HSet(ctx, key, value).Err()

	return err
}

func (c *RedisConfig) HGet(key string, field string) (string, error) {
	value, err := c.client.HGet(ctx, key, field).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return value, nil
}
