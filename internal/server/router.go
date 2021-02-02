package server

import (
	"github.com/AleksandrAkhapkin/true-conf/internal/server/handlers"
	"github.com/labstack/echo"
)

func NewRouter(h *handlers.Handlers) *echo.Echo {

	e := echo.New()

	e.GET("/ping", h.Ping)

	return e
}
