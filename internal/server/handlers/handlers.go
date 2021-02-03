package handlers

import (
	"github.com/AleksandrAkhapkin/true-conf/internal/service"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Handlers struct {
	srv       *service.Service
	secretKey string
}

func NewHandlers(srv *service.Service) (*Handlers, error) {
	return &Handlers{
		srv: srv,
	}, nil
}

func (h *Handlers) Ping(c echo.Context) error {

	if _, err := c.Response().Writer.Write([]byte("pong")); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in Ping"))
		return err
	}
	return nil
}

func (h *Handlers) CreateUser(c echo.Context) error {

	req := c.Request()
	userName := req.FormValue("name")

	err := h.srv.CreateUser(userName)
	if err != nil {
		return err
	}

	return nil
}

//func (h *Handlers) GetUsers(c echo.Context) error {
//
//
//	users, err := h.srv.GetUsers()
//	if err != nil {
//		return err
//	}
//
//	if _,  err := c.Response().Writer.Write([]byte(users)); err != nil {
//		return err
//	}
//
//	return nil
//}
