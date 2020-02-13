package exception

import (
	"github.com/getsentry/sentry-go"
)

// Send an error to Sentry.
// This requires sentry endpoint configured into the config.yml file
func SentryReport(report error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: configuration.Exception.Sentry,
	}); err != nil {
		Log(err.Error())
	}

	sentry.CaptureException(report)
}
