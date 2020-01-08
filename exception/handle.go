package exception

import (
	"fmt"
	"os"
)

// Generic method to handle errors.
// You can customize this method to implement your error handling.
// Es.: You can implement "Sentry" or other error tracking system
func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
