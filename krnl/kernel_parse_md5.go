package krnl

func (receiver *Kernel) parse_md5(src string) error {
	if nil == receiver {
		return errNilReceiver
	}

	ln := len(src)

	switch {
	case 32 == ln:
		return receiver.parse_hexadecimal(src)
	case 22 == ln || 23 == ln || 24 == ln:
		return receiver.parse_base64(src)
	}

	return errorf("unsupported format for MD5 digest — %q", src)
}
