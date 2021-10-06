package krnl

import (
	"encoding/ascii85"
	"strings"
	"io"
	"io/ioutil"
)

func (receiver *Kernel) parse_ascii85(src string) error {
	if nil == receiver {
		return errNilReceiver
	}

	var reader io.Reader = strings.NewReader(src)
	if nil == reader {
		return errInternalError
	}

	var decoder io.Reader = ascii85.NewDecoder(reader)
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
