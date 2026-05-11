package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

func NewHandler() *SBaiHandler {
	return &SBaiHandler{
		mutex: &sync.Mutex{},
	}
}

func (h *SBaiHandler) Handle(context context.Context, record slog.Record) error {
	var color, emoji string

	switch record.Level {
	case slog.LevelDebug:
		color = cGray
		emoji = eDebug
	case slog.LevelInfo:
		color = cBlue
		emoji = eInfo
	case slog.LevelWarn:
		color = cYellow
		emoji = eWarn
	case slog.LevelError:
		color = cRed
		emoji = eError
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	fmt.Fprintf(
		os.Stdout,
		"%s[%s] %s [%s]\t%s%s\t",
		color,
		record.Time.Format(timeStr),
		emoji,
		record.Level,
		record.Message,
		cReset,
	)

	record.Attrs(func(attr slog.Attr) bool {
		fmt.Fprintf(
			os.Stdout,
			" %s%s=%v%s",
			cGreen,
			attr.Key,
			attr.Value.Any(),
			cReset,
		)

		return true
	})

	fmt.Fprintln(os.Stdout)

	return nil
}

func (h *SBaiHandler) Enabled(context context.Context, level slog.Level) bool {
	return true
}
func (h *SBaiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SBaiHandler{
		mutex: h.mutex,
	}
}
func (h *SBaiHandler) WithGroup(name string) slog.Handler {
	return &SBaiHandler{
		mutex: h.mutex,
	}
}
