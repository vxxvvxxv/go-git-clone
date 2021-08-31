// Copied from: https://github.com/go-git/go-git/blob/master/_examples/common.go#L19

package common

import (
	"fmt"
	"os"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error, errMessage string) {
	if err == nil {
		return
	}

	fmt.Printf("%s\n\x1b[31;1m%s\x1b[0m\n", errMessage, fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
