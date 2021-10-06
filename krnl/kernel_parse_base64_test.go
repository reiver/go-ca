package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"

	"testing"
)

func TestKernel_Parse_base64_success(t *testing.T) {

	tests := []struct{
		Src string
		Expected []byte
	}{
		{
			Src: "",
			Expected: []byte{},
		},



		{
			Src: "SGVsbG8gd29ybGQh",
			Expected: []byte("Hello world!"),
		},



		{
			Src: "YXBwbGU=",
			Expected: []byte("apple"),
		},
		{
			Src: "YmFuYW5h",
			Expected: []byte("banana"),
		},
		{
			Src: "Y2hlcnJ5",
			Expected: []byte("cherry"),
		},



		{
			Src: "QQ==",
			Expected: []byte("A"),
		},
		{
			Src: "QUI=",
			Expected: []byte("AB"),
		},
		{
			Src: "QUJD",
			Expected: []byte("ABC"),
		},
		{
			Src: "QUJDRA==",
			Expected: []byte("ABCD"),
		},
		{
			Src: "QUJDREU=",
			Expected: []byte("ABCDE"),
		},
		{
			Src: "QUJDREVG",
			Expected: []byte("ABCDEF"),
		},
		{
			Src: "QUJDREVGRw==",
			Expected: []byte("ABCDEFG"),
		},
		{
			Src: "QUJDREVGR0g=",
			Expected: []byte("ABCDEFGH"),
		},
		{
			Src: "QUJDREVGR0hJ",
			Expected: []byte("ABCDEFGHI"),
		},
		{
			Src: "QUJDREVGR0hJSg==",
			Expected: []byte("ABCDEFGHIJ"),
		},



		{
			Src: "//7//v/+",
			Expected: []byte{0xff,0xfe,0xff,0xfe,0xff,0xfe},
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.Parse(test.Src, "base64")
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("SOURCE: %q", test.Src)
				continue
			}
		}

		{
			var expected []byte = test.Expected
			var actual   []byte = kernel.Bytes()

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual value of the kernel as a []byte is not what was expected.", testNumber)
				t.Logf("EXPECTED: #%v", expected)
				t.Logf("ACTUAL:   #%v", actual)
				t.Logf("SOURCE: %q", test.Src)
				continue
			}
		}
	}
}

func TestKernel_Parse_base64_failure(t *testing.T) {

	tests := []struct{
		Src string
	}{
		{
			Src: "<[+]>",
		},



		{
			Src: "!@#$%^&*()_+",
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		err := kernel.Parse(test.Src, "base64")

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("SOURCE: %q", test.Src)
			t.Logf("ERROR: %v", err)
			continue
		}
	}
}
