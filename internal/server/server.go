package server

import (
	"fmt"
	"github.com/AleksandrAkhapkin/true-conf/internal/server/handlers"
	"github.com/AleksandrAkhapkin/true-conf/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
)

func StartServer(handlers *handlers.Handlers, port string) {

	echo := NewRouter(handlers)
	logger.LogInfo("Restart service")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), echo); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
	}

}
