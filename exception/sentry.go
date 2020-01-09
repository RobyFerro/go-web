package exception

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

func SentryReport(report error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: configuration.Exception.Sentry,
	}); err != nil {
		fmt.Println(err.Error())
	}

	sentry.CaptureException(report)
}
