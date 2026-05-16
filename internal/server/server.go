package server

import (
	"fmt"
	"net/http"

	"go-links/internal/api/Url/routes"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	port int
	db   *gorm.DB
}

func NewServer(db *gorm.DB) *http.Server {
	e := echo.New()

	s := &Server{
		echo: e,
		port: 8080,
		db:   db,
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

	// Link Routes
	routes.RegisterLinkRoutes(s.echo, v1, s.db)
}
