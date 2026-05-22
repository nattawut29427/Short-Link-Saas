package middleware

import (
	"fmt"
	"net/http"
	"time"

	echoMiddleware "github.com/labstack/echo/v5/middleware"
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
)


func RateLimit() echo.MiddlewareFunc {
	config := echoMiddleware.RateLimiterConfig{
		Skipper: echoMiddleware.DefaultSkipper,
		Store: echoMiddleware.NewRateLimiterMemoryStoreWithConfig(
			echoMiddleware.RateLimiterMemoryStoreConfig{
				Rate:      10,                // requests per second
				Burst:     10,                // max burst capacity
				ExpiresIn: 3 * time.Minute,   // expires after 3 minutes of inactivity
			},
		),
		IdentifierExtractor: func(c *echo.Context) (string, error) {
			return c.RealIP(), nil
		},
		DenyHandler: func(c *echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "too many requests, please try again later",
			})
		},
	}

	return echoMiddleware.RateLimiterWithConfig(config)
}

func RedisRateLimit(rdb *redis.Client, limit int, window time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			ip := c.RealIP()
		
			now := time.Now().UnixNano()
			windowNs := window.Nanoseconds()
			currentWindow := now / windowNs
			key := fmt.Sprintf("rate:%s:%d", ip, currentWindow)

			ctx := c.Request().Context()
			count, err := rdb.Incr(ctx, key).Result()
			if err != nil {
				return next(c)
			}

			if count == 1 {
				rdb.Expire(ctx, key, window)
			}

			if count > int64(limit) {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error": "too many requests, please try again later",
				})
			}

			return next(c)
		}
	}
}