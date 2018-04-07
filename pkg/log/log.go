package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

// Level refers to the level of logging
type Level string

func (l Level) parseLevel() (zerolog.Level, error) {
	level := strings.ToLower(string(l))

	switch level {
	case "panic":
		return zerolog.PanicLevel, nil
	case "fatal":
		return zerolog.FatalLevel, nil
	case "error":
		return zerolog.ErrorLevel, nil
	case "warn":
		return zerolog.WarnLevel, nil
	case "info":
		return zerolog.InfoLevel, nil
	case "debug":
		return zerolog.DebugLevel, nil
	}

	return 0, fmt.Errorf("wrong level")
}

// Logger is an interface that needs to be implemented in order to log.
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warningf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})

	With(key string, value string) Logger
	Set(level Level) error
}

type logger struct {
	logger zerolog.Logger
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.sourced().Debug().Msgf(format, args...)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.sourced().Info().Msgf(format, args...)
}

func (l logger) Warningf(format string, args ...interface{}) {
	l.sourced().Warn().Msgf(format, args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.sourced().Error().Msgf(format, args...)
}
func (l logger) Fatalf(format string, args ...interface{}) {
	l.sourced().Fatal().Msgf(format, args...)
}
func (l logger) Panicf(format string, args ...interface{}) {
	l.sourced().Panic().Msgf(format, args...)
}

func (l logger) With(key string, value string) Logger {
	return &logger{l.logger.With().Str(key, value).Logger()}
}

func (l *logger) Set(lv Level) error {
	parsed, err := lv.parseLevel()
	if err != nil {
		return err
	}
	l.logger = l.logger.Level(parsed)
	return nil
}

func (l logger) sourced() *zerolog.Logger {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	zl := l.logger.With().Str("src", fmt.Sprintf("%s:%d", file, line)).Logger()
	return &zl
}

var baseLogger = &logger{
	logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
}

// Base returns the base logger
func Base() Logger {
	return baseLogger
}
