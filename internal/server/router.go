package server

import (
	"github.com/AleksandrAkhapkin/true-conf/internal/server/handlers"
	"github.com/labstack/echo"
)

func NewRouter(h *handlers.Handlers) *echo.Echo {

	e := echo.New()

	e.GET("/ping", h.Ping)
	e.GET("/createuser", h.CreateUser)
	e.GET("/users", h.GetUsers)
	return e
}
