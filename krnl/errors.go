package krnl

import (
	"errors"
)

var (
	errInternalError = errors.New("krnl: internal error")
	errNilReceiver   = errors.New("krnl: nil receiver")
	errNothing       = errors.New("krnl: nothing")
)
