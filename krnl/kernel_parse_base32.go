package krnl

import (
	"encoding/base32"
	"io/ioutil"
	"strings"
)

func (receiver *Kernel) parse_base32(src string) error {
	if nil == receiver {
		return errNilReceiver
	}

	reader := strings.NewReader(src)
	if nil == reader {
		return errInternalError
	}

	decoder := base32.NewDecoder(base32.StdEncoding, reader)
	if nil == decoder {
		return errInternalError
	}

	p, err := ioutil.ReadAll(decoder)
	if nil != err {
		return err
	}

	return receiver.UnmarshalBinary(p)
}
