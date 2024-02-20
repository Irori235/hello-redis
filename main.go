package main

import (
	"context"
	"fmt"
	"os"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redis := goredis.NewClient(&goredis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	// Subscribe
	go func() {
		pubsub := redis.Subscribe(ctx, "messages")
		defer pubsub.Close()

		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println("Received message:", msg.Payload)
		}
	}()

	// Publish
	time.Sleep(2 * time.Second)
	go func() {
		err := redis.Publish(ctx, "messages", "Hello from go-redis!").Err()
		if err != nil {
			fmt.Println("Publish error:", err)
			return
		}
	}()

	// Wait
	time.Sleep(5 * time.Second)
}
