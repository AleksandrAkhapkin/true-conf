package handlers

import (
	"fmt"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

//Создать пользователя
func (h *Handlers) CreateUser(c echo.Context) error {

	req := c.Request()
	userName := req.FormValue("name")
	if userName == "" {
		if _, err := c.Response().Write([]byte("Вы не указали имя для пользователя")); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}

	id, err := h.srv.CreateUser(userName)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Пользователь успешно зарегистрирован, его ID: %d", id)
	if _, err := c.Response().Writer.Write([]byte(message)); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in GetUsers"))
		return err
	}

	return nil
}

//Получить пользователя
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

	user, err := h.srv.GetUser(int32(id))
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

//Редактировать имя пользователя
func (h *Handlers) PutUserByID(c echo.Context) error {

	req := c.Request()

	name := req.FormValue("name")
	if name == "" {
		if _, err := c.Response().Write([]byte("Вы не указали новое имя для пользователя")); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}

	if err = h.srv.PutUser(int32(id), name); err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
			return err
		}
		return nil
	}

	if _, err := c.Response().Write([]byte("Имя пользователя успешно изменено")); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in PutUserByID"))
		return err
	}

	return nil
}

//Удалить пользователя
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

	if err = h.srv.DeleteUser(int32(id)); err != nil {
		if _, err := c.Response().Write([]byte(err.Error())); err != nil {
			logger.LogError(errors.Wrap(err, "err with responseWriter in DeleteUserByID"))
			return err
		}
		return nil
	}

	message := fmt.Sprintf("Пользователь с ID: %d успешно удален", id)
	if _, err := c.Response().Write([]byte(message)); err != nil {
		logger.LogError(errors.Wrap(err, "err with responseWriter in DeleteUserByID"))
		return err
	}

	return nil
}

//Получить всех пользователей
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
