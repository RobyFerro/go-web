package exception

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

// Generic method to handle errors.
// You can customize this method to implement your error handling.
// Es.: You can implement "Sentry" or other error tracking system
func ProcessError(err error) {
	SentryReport(err)
}

func SentryReport(report error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://4b98c76e810948eca928a81fd9e4e6c4@sentry.io/1876468",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	sentry.CaptureException(report)
}
