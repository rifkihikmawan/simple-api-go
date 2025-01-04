package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) handleHello(c echo.Context) error {
	name := c.Param("name")

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Hello, %s!", name),
	})
}
