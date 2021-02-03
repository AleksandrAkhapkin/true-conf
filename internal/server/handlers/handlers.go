package handlers

import (
	"github.com/AleksandrAkhapkin/true-conf/internal/service"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
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

	if _, err := c.Response().Writer.Write([]byte("Пользователь успешно зарегистрирован")); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in GetUsers"))
		return err
	}

	return nil
}

func (h *Handlers) GetUsers(c echo.Context) error {

	users, err := h.srv.GetUsers()
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in GetUsers"))
			return err
		}
		return nil
	}

	if err := c.JSON(http.StatusOK, users); err != nil {
		logger.LogError(errors.Wrap(err, "err with c.JSON in GetUsers"))
		return err
	}

	return nil
}

func (h *Handlers) GetUserByID(c echo.Context) error {

	req := c.Request()

	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in GetUserByID"))
			return err
		}
		return nil
	}

	user, err := h.srv.GetUser(id)
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in GetUserByID"))
			return err
		}
		return nil
	}

	if err := c.JSON(http.StatusOK, user); err != nil {
		logger.LogError(errors.Wrap(err, "err with c.JSON in GetUserByID"))
		return err
	}

	return nil
}

func (h *Handlers) PutUserByID(c echo.Context) error {

	req := c.Request()

	name := req.FormValue("name")
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}

	if err = h.srv.PutUser(id, name); err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}

	if _, err := c.Response().Write([]byte("Пользователь успешно изменен")); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
		return err
	}

	return nil
}

func (h *Handlers) DeleteUserByID(c echo.Context) error {

	req := c.Request()

	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in DeleteUserByID"))
			return err
		}
		return nil
	}

	if err = h.srv.DeleteUser(id); err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in DeleteUserByID"))
			return err
		}
		return nil
	}

	if _, err := c.Response().Write([]byte("Пользователь успешно удален")); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in DeleteUserByID"))
		return err
	}

	return nil
}
