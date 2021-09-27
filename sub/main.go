package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	subscriber := redisClient.Subscribe(ctx, "user-data-topi")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from "+msg.Channel+" channel.", msg.Payload)
		fmt.Printf("%+v\n", user)
	}
}
