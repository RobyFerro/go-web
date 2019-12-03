package exception

import (
	"fmt"
	"os"
)

func ProcessError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
