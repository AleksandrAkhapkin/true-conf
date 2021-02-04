package service

import (
	"fmt"
	"github.com/AleksandrAkhapkin/true-conf/internal/types"
	"sync/atomic"
)

func (s *Service) CreateUser(userName string) (int32, error) {

	atomic.AddInt32(&s.lastID, 1)

	user := types.User{
		ID:   s.lastID,
		Name: userName,
	}

	s.users.Users = append(s.users.Users, user)

	return s.lastID, nil
}

func (s *Service) GetUsers() (*types.Users, error) {

	if s.users.Users == nil {
		return nil, fmt.Errorf("Пользователей еще нету")
	}
	return s.users, nil
}

func (s *Service) GetUser(id int32) (*types.User, error) {

	if s.users.Users == nil {
		return nil, fmt.Errorf("Пользователей еще нету")
	}

	for _, v := range s.users.Users {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Пользователь c ID: %d не найден", id)
}

func (s *Service) PutUser(id int32, name string) error {

	if s.users.Users == nil {
		return fmt.Errorf("Пользователей еще нету")
	}

	for i, v := range s.users.Users {
		if v.ID == id {
			s.users.Users[i].Name = name
			return nil
		}
	}

	return fmt.Errorf("Пользователь c ID: %d не найден", id)
}

func (s *Service) DeleteUser(id int32) error {

	if s.users.Users == nil {
		return fmt.Errorf("Пользователей еще нету")
	}

	usersTmp := &types.Users{}
	del := false
	for _, v := range s.users.Users {
		if v.ID == id {
			del = true
			continue
		}
		usersTmp.Users = append(usersTmp.Users, v)
	}

	if !del {
		return fmt.Errorf("Пользователь с ID: %d не найден", id)
	}
	s.users = usersTmp

	return nil
}
