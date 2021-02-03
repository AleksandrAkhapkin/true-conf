package service

import (
	"fmt"
	"github.com/AleksandrAkhapkin/true-conf/internal/types"
)

func (s *Service) CreateUser(userName string) error {

	//todo mutex
	s.lastID++

	user := types.User{
		ID:   s.lastID,
		Name: userName,
	}

	s.users.Users = append(s.users.Users, user)

	return nil
}

func (s *Service) GetUsers() (*types.Users, error) {

	if s.users.Users == nil {
		return nil, fmt.Errorf("Пользователей еще нету")
	}
	return s.users, nil
}
