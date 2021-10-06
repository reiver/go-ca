package krnl

import (
	"encoding/base64"
	"strings"
	"io"
	"io/ioutil"
)

func (receiver *Kernel) parse_base64url(src string) error {
	if nil == receiver {
		return errNilReceiver
	}

	var reader io.Reader = strings.NewReader(src)
	if nil == reader {
		return errInternalError
	}

	var decoder io.Reader = base64.NewDecoder(base64.RawURLEncoding, reader)
	if nil == decoder {
		return errInternalError
	}

	p, err :=  ioutil.ReadAll(decoder)
	if nil != err {
		return err
	}

	if err := receiver.UnmarshalBinary(p); nil != err {
		return err
	}

	return nil
}
