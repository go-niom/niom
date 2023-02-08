package logger

const Logger = `package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"{{ .ModuleName}}/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var logger zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		appCfg := config.AppCfg()
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(appCfg.LogLevel)
		if err != nil {
			logLevel = int(zerolog.InfoLevel) // default to INFO
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if appCfg.AppEnv != "development" {
			fileLogger := &lumberjack.Logger{
				Filename:   "wikipedia-demo.log",
				MaxSize:    5, //
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		logger = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Logger()
	})

	return logger
}

func Components(name, method string) zerolog.Logger {
	return logger.With().Str("name", name).Str("method", method).Logger()
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Error(msg string, err error) {
	logger.Error().Err(err).Msg(msg)
}

func Warn(msg string) {
	logger.Warn().Msg(msg)
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}
`
