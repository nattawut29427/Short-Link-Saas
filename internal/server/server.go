package server

import (
	"fmt"
	"net/http"

	loginRoutes "go-links/internal/api/auth/login/routes"
	authRoutes "go-links/internal/api/auth/register/routes"
	urlRoutes "go-links/internal/api/url/routes"
	"go-links/internal/middleware"

	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	port int
	db   *gorm.DB
	rdb  *redis.Client
	jwt  string
}

func NewServer(port int, db *gorm.DB, rdb *redis.Client, jwt string) *http.Server {
	e := echo.New()

	s := &Server{
		echo: e,
		port: port,
		db:   db,
		rdb:  rdb,
		jwt:  jwt,
	}

	e.Use(middleware.Recover())

	s.MainRoutes()

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.echo,
	}
}

func (s *Server) MainRoutes() {
	// Root Group
	v1 := s.echo.Group("/v1")

	// Health Check
	v1.GET("/health", func(c *echo.Context) error {
		sqlDB, err := s.db.DB()
		dbStatus := "ok"
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "unavailable"
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
			"db":     dbStatus,
		})
	})

	v1.GET("/panic", func(c *echo.Context) error {
	panic("test recover middleware")
})
	
	RateLimitGroup := v1.Group("", middleware.RateLimit())

	authRoutes.RegisterAuthRoutes(s.echo, RateLimitGroup, s.db, s.rdb, s.jwt)
	loginRoutes.LoginRoutes(s.echo, RateLimitGroup, s.db, s.rdb, s.jwt)

	protected := v1.Group("", middleware.JWTAuth(s.jwt))

	urlRoutes.RegisterLinkRoutes(s.echo, protected, s.db, s.rdb)
	
}

