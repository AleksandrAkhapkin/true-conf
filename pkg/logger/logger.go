package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var logger = zerolog.New(os.Stdout)

func LogError(err error) {
	logger.Err(err).Time("time", time.Now()).Send()
}

func LogInfo(msg string) {
	logger.Info().Time("time", time.Now()).Msg(msg)
}

func LogFatal(err error) {
	logger.Fatal().Err(err).Time("time", time.Now()).Send()
}
