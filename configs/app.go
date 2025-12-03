package configs

import "github.com/gofiber/fiber/v2"

var (
	Port        = ":10086"
	FiberConfig = fiber.Config{
		DisableStartupMessage: false,
	}
)
