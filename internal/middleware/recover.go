// package middleware

// import (
// 	"github.com/labstack/echo/v5"
// 	echoMiddleware "github.com/labstack/echo/v5/middleware"
// )

// func Recover() echo.MiddlewareFunc {

// 	config := echoMiddleware.RecoverConfig{

// 		Skipper:           echoMiddleware.DefaultSkipper,
// 		StackSize:         4 << 10, 
// 		DisableStackAll:   false,
// 		DisablePrintStack: false,
// 	}

// 	return  echoMiddleware.RecoverWithConfig(config)
// }
package middleware

import (
	"log/slog"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v5"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					buf := make([]byte, 4<<10)
					n := runtime.Stack(buf, false)
					slog.Error("panic recovered",
						"error",  r,
						"path",   c.Request().URL.Path,
						"method", c.Request().Method,
						"stack",  string(buf[:n]),
					)
						c.JSON(http.StatusInternalServerError, map[string]string{
						"message": "Internal Server Error",
					})
				}
			}()
			return next(c)
		}
	}
}