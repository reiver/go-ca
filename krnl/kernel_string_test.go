package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"testing"
)

func TestKernel_String_success(t *testing.T) {

	tests := []struct{
		Src interface{}
		Expected string
	}{
		{
			Src: []byte{},
			Expected: "",
		},



		{
			Src: "apple",
			Expected: "YXBwbGU",
		},
		{
			Src: "banana",
			Expected: "YmFuYW5h",
		},
		{
			Src: "cherry",
			Expected: "Y2hlcnJ5",
		},



		{
			Src: []byte{
				0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
				0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
				0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
				0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
				0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
				0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
				0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
				0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
			},
			Expected: "ld7McvClCuTZ1TeOGyJSWHz8cZd-QyksjxuEZIJIUJ8bwYvG8LDQuGBqZD7_YdYRroTm-9SiaDFlcGvW_UizNA",
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.Scan(test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		{
			var expected string = test.Expected
			var actual   string = kernel.String()

			if expected != actual {
				t.Errorf("For test #%d, actual result of .String() is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}

		}
	}
}

func TestKernel_String_failure(t *testing.T) {

	var kernel krnl.Kernel

	actual := kernel.String()

	if expected := "<error:not-loaded>"; expected != actual {
		t.Errorf("The actual value is not what was expectd.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ACTUAL:   %q", actual)
		return
	}
}
