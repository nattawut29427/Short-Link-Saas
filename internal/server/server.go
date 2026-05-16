package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	port int
	db   *gorm.DB
}

func NewServer(port int, db *gorm.DB) *http.Server {
	e := echo.New()

	s := &Server{
		echo: e,
		port: port,
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
	v1.GET("/health", func(c echo.Context) error {
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

	// Address Routes
	// routes.RegisterAddressRoutes(v1, s.db)
}
