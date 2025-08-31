package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewSimpleLogger(serviceName, environment, level string) zerolog.Logger {
	var logLevel zerolog.Level

	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	// readable console in dev, json in prod
	var writer zerolog.ConsoleWriter
	if environment == "production" {
		writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	}

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return zerolog.New(writer).
		Level(logLevel).
		With().
		Timestamp().
		Str("service", serviceName).
		Str("environment", environment).
		Logger()
}
