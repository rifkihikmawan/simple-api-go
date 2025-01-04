package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) handleHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
