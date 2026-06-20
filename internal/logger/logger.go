package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/lumberjack.v3"
)

type Config struct {
	Level    slog.Level
	FilePath string
	Pretty   bool
}

func New(cfg Config) (*slog.Logger, error) {
	var writers []io.Writer
	writers = append(writers, os.Stderr)

	if cfg.FilePath != "" {
		rotator, err := lumberjack.New(
			lumberjack.WithFileName(cfg.FilePath),
			lumberjack.WithMaxBytes(10*lumberjack.MB),
			lumberjack.WithMaxBackups(5),
			lumberjack.WithMaxDays(30),
			lumberjack.WithCompress(),
		)
		if err != nil {
			return nil, err
		}
		writers = append(writers, rotator)
	}

	w := io.MultiWriter(writers...)

	opts := &slog.HandlerOptions{Level: cfg.Level}

	var handler slog.Handler
	if cfg.Pretty {
		handler = slog.NewTextHandler(w, opts)
	} else {
		handler = slog.NewJSONHandler(w, opts)
	}

	return slog.New(handler), nil
}
