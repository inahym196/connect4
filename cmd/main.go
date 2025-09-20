package main

import (
	"log/slog"

	"github.com/inahym196/connect4/internal"
)

func main() {
	slog.SetDefault(internal.NewLogger())
	slog.Debug("hello")
}
