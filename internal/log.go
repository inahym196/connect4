package internal

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

func NewLogger() *slog.Logger {
	wd, _ := os.Getwd()
	ops := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				return slog.String(a.Key, time.Now().Format("15:04:05"))
			case slog.SourceKey:
				if src, ok := a.Value.Any().(*slog.Source); ok {
					path := strings.Replace(src.File, wd, ".", 1)
					function := strings.Split(src.Function, "/")
					return slog.String(a.Key, fmt.Sprintf("%s:%s:%d", path, function[len(function)-1], src.Line))
				}
			}
			return a
		},
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &ops))
}
