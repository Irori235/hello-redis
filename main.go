package main

import (
	"context"
	"fmt"
	"os"

	goredis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redis := goredis.NewClient(&goredis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	if err := redis.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	defer redis.Close()

	cmd := redis.Set(ctx, "key", "value", 0)

	if err := cmd.Err(); err != nil {
		panic(err)
	}

	val, err := redis.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)
}
