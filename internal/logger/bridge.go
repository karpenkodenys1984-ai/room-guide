package logger

import (
	"log/slog"
	"time"
)

type FrontendLogEntry struct {
	Level     string         `json:"level"`
	Message   string         `json:"message"`
	Component string         `json:"component"`
	Extra     map[string]any `json:"extra,omitempty"`
	Timestamp string         `json:"timestamp"`
}

type Bridge struct {
	log *slog.Logger
}

func NewBridge(log *slog.Logger) *Bridge {
	return &Bridge{log: log}
}

func (b *Bridge) LogFromFrontend(entry FrontendLogEntry) {
	t, err := time.Parse(time.RFC3339, entry.Timestamp)
	if err != nil {
		t = time.Now() // fallback если не распарсилось
	}

	attrs := []any{
		slog.String("source", "frontend"),
		slog.String("component", entry.Component),
		slog.Time("client_time", t),
	}

	for k, v := range entry.Extra {
		attrs = append(attrs, slog.Any(k, v))
	}

	switch entry.Level {
	case "error":
		b.log.Error(entry.Message, attrs...)
	case "warn":
		b.log.Warn(entry.Message, attrs...)
	case "debug":
		b.log.Debug(entry.Message, attrs...)
	default:
		b.log.Info(entry.Message, attrs...)
	}
}
