package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var Logger *slog.Logger

func Set(logger *slog.Logger) {
	Logger = logger
}

func UseDefault() {
	logHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "2006-01-02 15:04:05.000",
		ReplaceAttr: func(groups []string, attribute slog.Attr) slog.Attr {
			if attribute.Key == slog.LevelKey {
				level, ok := attribute.Value.Any().(slog.Level)
				if ok {
					switch level {
					case slog.LevelDebug:
						return tint.Attr(4, slog.String(attribute.Key, "DBG"))
					case slog.LevelInfo:
						return tint.Attr(2, slog.String(attribute.Key, "INF"))
					case slog.LevelWarn:
						return tint.Attr(3, slog.String(attribute.Key, "WRN"))
					case slog.LevelError:
						return tint.Attr(1, slog.String(attribute.Key, "ERR"))
					}
				}
			}

			return attribute
		},
	})

	Logger = slog.New(logHandler)
}

func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	Logger.Error(msg, args...)
}
