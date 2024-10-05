package utils

import (
	"log/slog"
	"os"
	"time"

	slogmulti "github.com/samber/slog-multi"
	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"

	"github.com/joho/godotenv"
)

func slogLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	return logger
}

func sentryLogger() *slog.Logger {
	err := godotenv.Load()
	if err != nil {
		sentry.CaptureException(err)
	}
	dsn := os.Getenv("SENTRY_DSN")

	err = sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		EnableTracing: false,
	})
	if err != nil {
		slog.Error("sentry.Init: " + err.Error())
	}

	defer sentry.Flush(2 * time.Second)

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())

	return logger
}

func Logger() *slog.Logger {
	slogLogger := slogLogger()
	sentryLogger := sentryLogger()

	logger := slog.New(
		slogmulti.Fanout(
			slogLogger.Handler(),
			sentryLogger.Handler(),
		),
	)

	return logger
}
