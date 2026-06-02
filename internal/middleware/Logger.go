package middleware 

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ZLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()

			traceID := uuid.New().String()
			
			c.Response().Header().Set("X-Request-ID", traceID)
			reqLogger := log.With().Str("trace_id", traceID).Logger()
			c.Set("logger", &reqLogger)
			
			err := next(c)

			_, status := echo.ResolveResponseStatus(c.Response(), err)
			req := c.Request()

	
			var event *zerolog.Event
			if err != nil || status >= 500 {
				event = reqLogger.Error()
			} else if status >= 400 {
				event = reqLogger.Warn()
			} else {
				event = reqLogger.Info()
			}

			event.
				Str("method", req.Method).
				Str("uri", req.RequestURI).
				Int("status", status).
				Dur("latency", time.Since(start)).
				Str("ip", c.RealIP()).
				Str("user_agent", req.UserAgent())

			if err != nil {
				event.Err(err)
			}

			event.Msg("HTTP request")

			return err
		}
	}
}