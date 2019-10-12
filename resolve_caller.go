package snapshot

import (
	"errors"
	"runtime"
	"strings"
)

const maxCallerSkip int = 50

func getCallerFilename() string {
	for i := 0; i < maxCallerSkip; i++ {
		_, fn, _, _ := runtime.Caller(i)
		if strings.HasSuffix(fn, "_test.go") || strings.HasSuffix(fn, "_tests.go") {
			return fn
		}
	}
	panic(errors.New("Unable to find a test caller, make sure you're importing from a test file"))
}
