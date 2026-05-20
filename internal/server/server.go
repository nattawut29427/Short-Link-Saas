package server

import (
	"fmt"
	"net/http"

	authRoutes "go-links/internal/api/auth/register/routes"
	urlRoutes "go-links/internal/api/url/routes"
	loginRoutes "go-links/internal/api/auth/login/routes"

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

	s.RegisterRoutes()

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.echo,
	}
}

func (s *Server) RegisterRoutes() {
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

	urlRoutes.RegisterLinkRoutes(s.echo, v1, s.db, s.rdb)
	authRoutes.RegisterAuthRoutes(s.echo, v1, s.db, s.rdb, s.jwt)
	loginRoutes.LoginRoutes(s.echo, v1, s.db, s.rdb, s.jwt)
}
