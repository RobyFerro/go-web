package exception

import (
	"github.com/getsentry/sentry-go"
)

func SentryReport(report error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: configuration.Exception.Sentry,
	}); err != nil {
		Log(err.Error())
	}

	sentry.CaptureException(report)
}
