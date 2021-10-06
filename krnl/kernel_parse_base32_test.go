package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"

	"testing"
)

func TestKernel_Parse_base32_success(t *testing.T) {

	tests := []struct{
		Src string
		Expected []byte
	}{
		{
			Src: "",
			Expected: []byte{},
		},



		{
			Src: "JBSWY3DPEB3W64TMMQQQ====",
			Expected: []byte("Hello world!"),
		},



		{
			Src: "MFYHA3DF",
			Expected: []byte("apple"),
		},
		{
			Src: "MJQW4YLOME======",
			Expected: []byte("banana"),
		},
		{
			Src: "MNUGK4TSPE======",
			Expected: []byte("cherry"),
		},



		{
			Src: "IE======",
			Expected: []byte("A"),
		},
		{
			Src: "IFBA====",
			Expected: []byte("AB"),
		},
		{
			Src: "IFBEG===",
			Expected: []byte("ABC"),
		},
		{
			Src: "IFBEGRA=",
			Expected: []byte("ABCD"),
		},
		{
			Src: "IFBEGRCF",
			Expected: []byte("ABCDE"),
		},
		{
			Src: "IFBEGRCFIY======",
			Expected: []byte("ABCDEF"),
		},
		{
			Src: "IFBEGRCFIZDQ====",
			Expected: []byte("ABCDEFG"),
		},
		{
			Src: "IFBEGRCFIZDUQ===",
			Expected: []byte("ABCDEFGH"),
		},
		{
			Src: "IFBEGRCFIZDUQSI=",
			Expected: []byte("ABCDEFGHI"),
		},
		{
			Src: "IFBEGRCFIZDUQSKK",
			Expected: []byte("ABCDEFGHIJ"),
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.Parse(test.Src, "base32")
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

func TestKernel_Parse_base32_failure(t *testing.T) {

	tests := []struct{
		Src string
	}{
		{
			Src: "<([+])>",
		},



		{
			Src: "!@#$%^&*()_+",
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		err := kernel.Parse(test.Src, "base32")

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("SOURCE: %q", test.Src)
			t.Logf("ERROR: %v", err)
			continue
		}
	}
}
