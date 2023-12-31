package logger

import (
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/dusted-go/logging/prettylog"
)

func NewSlogHandler() slog.Logger {
	hostname, err := os.Hostname()
    if err != nil {
        hostname = ""
        slog.Error("could not get hostname")
    }
	prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
		Level:       slog.LevelDebug,
		AddSource:   true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "UTCTime"
				a.Value = slog.TimeValue(time.Now().Local())
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
        slog.Group("app_details",
            slog.Int("pid", os.Getpid()),
            slog.String("hostname", hostname),
            slog.String("go_version", runtime.Version()),
        ),
    })
    logger := slog.New(prettyHandler)
    slog.SetDefault(logger)

    return *logger
}
