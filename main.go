package main

import (
	"log/slog"

	"github.com/baturax/gutils/pkg/logger"
)

func main() {
	logger := slog.New(logger.NewHandler())
	slog.SetDefault(logger)

	slog.Debug("Ts is a Debug")
	slog.Info("Ts is an Info")
	slog.Warn("Ts is a Warn", "67:", "bai")
	slog.Error("Ts is an Error", "67:", "bai")
}
