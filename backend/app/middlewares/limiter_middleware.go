package middlewares

import (
	"time"

	"github.com/go-raptor/raptor"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewLimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Method() != "POST"
		},
		Max:        5,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return raptor.NewErrorTooManyRequests("Previše zahtjeva")
		},
	})
}
