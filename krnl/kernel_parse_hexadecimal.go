package krnl

import (
	"encoding/hex"
	"strings"
)

func (receiver *Kernel) parse_hexadecimal(src string) error {
	if nil == receiver {
		return errNilReceiver
	}

	{
		const prefix = "0x"

		if strings.HasPrefix(src, prefix) && len(prefix) < len(src) {
			src = src[len(prefix):]
		}
	}

	p, err := hex.DecodeString(src)
	if nil != err {
		return err
	}

	return receiver.UnmarshalBinary(p)
}
