package main

import (
	"flag"
	"fmt"
	"github.com/AleksandrAkhapkin/true-conf/internal/server"
	"github.com/AleksandrAkhapkin/true-conf/internal/server/handlers"
	"github.com/AleksandrAkhapkin/true-conf/internal/service"
	"github.com/AleksandrAkhapkin/true-conf/pkg/infrastructure"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
)

//Вам нужно написать маленькое приложение на ECHO framework предоставляющие REST API по работе с сущностью User.
//
//REST API должно удовлетворять следующие возможности:
//• Добавление User //TODO OK
//• Получение списка User
//• Получение User по Id
//• Редактирование User по Id
//• Удаление User по Id
//
//REST API должно работать с форматом данных JSON.
//
//Сущность User должно состоять минимум из следующих полей:
//• Идентификатор пользователя //TODO OK
//• Отображаемое имя //TODO OK
//
//Вы можете использовать дополнительные поля, если считаете нужным.
//
//В качестве хранилища данных нужно использовать файл в формате JSON.

func main() {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		err := fmt.Errorf("PANIC:'%v'\nRecovered in: %s", r, infrastructure.IdentifyPanic())
		logger.LogError(err)
	}()

	port := new(string)
	flag.StringVar(port, "server-port", "8080", "path to yaml config")
	flag.Parse()

	srv, err := service.NewService()
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewService"))
	}

	handls, err := handlers.NewHandlers(srv)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "with NewHandlers"))
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-sigChan
		logger.LogInfo("Finish service")
		srv.Close()
		os.Exit(0)
	}()

	server.StartServer(handls, *port)

}
