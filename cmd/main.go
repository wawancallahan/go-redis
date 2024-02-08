package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/wawancallahan/go-redis/internal/config"
)

func main() {
	redisClient := config.NewRedisConfig()

	// redisValue := "VALUES"
	jsonByte, err := json.Marshal(map[string]string{
		"data": "success",
	})

	redisValue := string(jsonByte)

	if err != nil {
		panic(err)
	}

	err = redisClient.Set("TEST", redisValue, 1*time.Hour)

	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get("TEST")

	if err != nil {
		panic(err)
	}

	var result map[string]string

	if err = json.Unmarshal([]byte(val), &result); err != nil {
		panic(err)
	}

	log.Println("REDIS VALUE =", val, result)

	err = redisClient.Delete("TEST")

	if err != nil {
		panic(err)
	}

	val2, err := redisClient.Get("TEST")

	if err != nil {
		panic(err)
	}

	log.Println("REDIS VALUE AFTER DELETED =", val2)
}
