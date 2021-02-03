package server

import (
	"github.com/AleksandrAkhapkin/true-conf/internal/server/handlers"
	"github.com/labstack/echo"
)

func NewRouter(h *handlers.Handlers) *echo.Echo {

	e := echo.New()

	e.GET("/ping", h.Ping)

	e.POST("/user", h.CreateUser)
	e.GET("/user", h.GetUserByID)
	e.DELETE("/user", h.DeleteUserByID)
	e.PUT("/user", h.PutUserByID)

	e.GET("/users", h.GetUsers)

	return e
}
