package service

import (
	"bufio"
	"encoding/json"
	"github.com/AleksandrAkhapkin/true-conf/internal/types"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/pkg/errors"
	"io"
	"os"
	"sort"
)

type Service struct {
	file   *os.File
	lastID int
	users  *types.Users
}

func NewService() (*Service, error) {

	lastID := 0
	file, err := openOrCreate()
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)
	users := &types.Users{}

	if err := json.NewDecoder(r).Decode(&users); err != nil {
		if err == io.EOF {
			return &Service{
				lastID: lastID,
				file:   file,
				users:  users,
			}, nil
		}
		logger.LogError(errors.Wrap(err, "err with Decode in NewService"))
		return nil, err
	}

	sort.Slice(users.Users, func(i, j int) (less bool) {
		return users.Users[i].ID > users.Users[j].ID
	})
	lastID = users.Users[0].ID

	return &Service{
		lastID: lastID,
		file:   file,
		users:  users,
	}, nil
}

func (s *Service) Close() {

	if s.users.Users != nil {
		b, err := json.Marshal(s.users)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err with Marshal in Close"))
			return
		}
		_, err = s.file.WriteAt(b, 0)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err with file.Write"))
			return
		}
	}
}

func openOrCreate() (*os.File, error) {

	file, err := os.OpenFile("users.json", os.O_RDWR, 0777)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.LogError(errors.Wrap(err, "err with Open in openOrCreate"))
			return nil, err
		}
		file, err = os.Create("users.json")
		if err != nil {
			logger.LogError(errors.Wrap(err, "err with Create  in openOrCreate"))
			return nil, err
		}
	}

	return file, nil
}
