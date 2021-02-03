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
	flag.StringVar(port, "port", "8080", "path to yaml config")
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
