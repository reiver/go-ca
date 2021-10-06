package krnl

import (
	"fmt"
)

func errorf(format string, a ...interface{}) error {
	format = fmt.Sprintf("ca: %s", format)

	return fmt.Errorf(format, a...)
}
