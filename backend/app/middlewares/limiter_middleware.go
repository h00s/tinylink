package middlewares

import (
	"time"

	"github.com/go-raptor/raptor"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewBurstLimiterMiddleware() fiber.Handler {
	return newLimiterMiddleware(5, 1*time.Minute)
}

func NewDailyLimiterMiddleware() fiber.Handler {
	return newLimiterMiddleware(50, 24*time.Hour)
}

func newLimiterMiddleware(max int, expiration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Method() != "POST"
		},
		Max:        max,
		Expiration: expiration,
		LimitReached: func(c *fiber.Ctx) error {
			return c.
				Status(fiber.StatusTooManyRequests).
				JSON(raptor.NewErrorTooManyRequests("Previše zahtjeva"))
		},
	})
}
