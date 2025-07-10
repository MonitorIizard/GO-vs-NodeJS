package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		start := time.Now()

		var sum int64 = 0
		for i := int64(0); i < 1_000_000_000; i++ {
			sum += i
		}

		duration := time.Since(start)
		message := fmt.Sprintf("\nGO Benchmark: %v\n", duration)
		message += fmt.Sprintf("Sum: %v\n\n", sum)
		return c.SendString(message)
	})

	app.Listen(":5555")
}
