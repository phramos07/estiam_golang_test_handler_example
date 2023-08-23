package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler funcation is sth like this: func(c echo.Context) error

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) IsAlive(c echo.Context) error {
	response := map[string]string{"message": "app is running"}
	return c.JSON(http.StatusOK, response)
}
