package krnl

import (
	"encoding/base64"
	"strings"
)

// String returns the value of the kernel in base64url format.
//
// So, for example, if the logical value of the kernel is:
//
//	[]byte{
//		0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
//		0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
//		0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
//		0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
//		0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
//		0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
//		0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
//		0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
//	}
//
// Then this .String() method would return:
//
//	ld7McvClCuTZ1TeOGyJSWHz8cZd-QyksjxuEZIJIUJ8bwYvG8LDQuGBqZD7_YdYRroTm-9SiaDFlcGvW_UizNA
//
//
// As an aside, one could parse this into a Kernel with:
//
//	var value string = "ld7McvClCuTZ1TeOGyJSWHz8cZd-QyksjxuEZIJIUJ8bwYvG8LDQuGBqZD7_YdYRroTm-9SiaDFlcGvW_UizNA"
//	
//	err := kernel.Parse(value, "base64url")
//
// String also makes krnl.Kernel fit the fmt.Stringer interface.
//
// That means it is likely one would indirectly call the .String() method when using fmt.Fprintf() fmt.Printf(), fmt.Sprintf().
func (receiver Kernel) String() string {
	const errorNotLoaded     = "<error:not-loaded>"
	const errorInternalError = "<error:internal-error>"

	if nothing() == receiver {
		return errorNotLoaded
	}

	var encoded strings.Builder
	{
		writeCloser := base64.NewEncoder(base64.RawURLEncoding, &encoded)
		if nil == writeCloser {
			return errorInternalError
		}

		n, err := writeCloser.Write(receiver.slice())
		if nil != err {
			return errorInternalError
		}
		if expected, actual := receiver.length, n; expected != actual {
			return errorInternalError
		}

		err = writeCloser.Close()
		if nil != err {
			return errorInternalError
		}
	}

	return encoded.String()
}
