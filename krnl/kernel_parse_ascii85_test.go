package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"

	"testing"
)

func TestKernel_Parse_ascii85_success(t *testing.T) {

	tests := []struct{
		Src string
		Expected []byte
	}{
		{
			Src: "",
			Expected: []byte{},
		},



		{
			Src: "87cURD]j7BEbo80",
			Expected: []byte("Hello world!"),
		},



		{
			Src: "@;p1%AH",
			Expected: []byte("apple"),
		},
		{
			Src: "@UX._DIE",
			Expected: []byte("banana"),
		},
		{
			Src: "@q]FoEd7",
			Expected: []byte("cherry"),
		},



		{
			Src: "5l",
			Expected: []byte("A"),
		},
		{
			Src: "5sb",
			Expected: []byte("AB"),
		},
		{
			Src: "5sdp",
			Expected: []byte("ABC"),
		},
		{
			Src: "5sdq,",
			Expected: []byte("ABCD"),
		},
		{
			Src: "5sdq,70",
			Expected: []byte("ABCDE"),
		},
		{
			Src: "5sdq,77I",
			Expected: []byte("ABCDEF"),
		},
		{
			Src: "5sdq,77Kc",
			Expected: []byte("ABCDEFG"),
		},
		{
			Src: "5sdq,77Kd<",
			Expected: []byte("ABCDEFGH"),
		},
		{
			Src: "5sdq,77Kd<8H",
			Expected: []byte("ABCDEFGHI"),
		},
		{
			Src: "5sdq,77Kd<8P/",
			Expected: []byte("ABCDEFGHIJ"),
		},



		{
			Src: "s8N&ts8D",
			Expected: []byte{0xff,0xfe,0xff,0xfe,0xff,0xfe},
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.Parse(test.Src, "ascii85")
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

func TestKernel_Parse_ascii85_failure(t *testing.T) {

	tests := []struct{
		Src string
	}{
		{
			Src: "~",
		},
		{
			Src: ">",
		},


		{
			Src: "-~",
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		err := kernel.Parse(test.Src, "ascii85")

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("SOURCE: %q", test.Src)
			t.Logf("ERROR: %v", err)
			continue
		}
	}
}
