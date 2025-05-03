package logging

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Level int
}

type logger struct {
	logger *slog.Logger
}

func Create(conf Config) (*logger, error) {
	var programLevel = new(slog.LevelVar) // Info by default
    level := conf.Level

	if level >= 4 {
		programLevel.Set(slog.LevelDebug)
	} else if level == 3 {
		programLevel.Set(slog.LevelInfo)
	} else if level == 2 {
		programLevel.Set(slog.LevelWarn)
	}

	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})

	l := slog.New(h)

	return &logger{logger: l}, nil
}

func (l logger) Debug(fmtStr string, vals ...any) {
	l.logger.Debug(format(fmtStr, vals))
}

func (l logger) Info(fmtStr string, vals ...any) {
	l.logger.Info(format(fmtStr, vals))
}

func (l logger) Warning(fmtStr string, vals ...any) {
	l.logger.Warn(format(fmtStr, vals))
}

func (l logger) Error(fmtStr string, vals ...any) {
	l.logger.Error(format(fmtStr, vals))
}

func format(fmtStr string, a []any) string {
	if len(a) == 0 {
		return fmtStr
	}

	return fmt.Sprintf(fmtStr, a...)
}
