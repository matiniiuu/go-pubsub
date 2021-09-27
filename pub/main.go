package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		if err := redisClient.Publish(ctx, "user-data-topi", payload).Err(); err != nil {
			panic(err)
		}

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
