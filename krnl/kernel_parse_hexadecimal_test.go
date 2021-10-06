package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"

	"testing"
)

func TestKernel_Parse_hexadecimal_success(t *testing.T) {

	tests := []struct{
		Src string
		Expected []byte
	}{
		{
			Src: "",
			Expected: []byte{},
		},



		{
			Src:              "fe",
			Expected: []byte{0xfe},
		},
		{
			Src:              "fe" +"dc",
			Expected: []byte{0xfe, 0xdc},
		},
		{
			Src:              "fe" +"dc" +"ba",
			Expected: []byte{0xfe, 0xdc, 0xba},
		},
		{
			Src:              "fe" +"dc" +"ba" +"98",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98},
		},
		{
			Src:              "fe" +"dc" +"ba" +"98" +"76",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76},
		},
		{
			Src:              "fe" +"dc" +"ba" +"98" +"76" +"54",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54},
		},
		{
			Src:              "fe" +"dc" +"ba" +"98" +"76" +"54" +"32",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32},
		},
		{
			Src:              "fe" +"dc" +"ba" +"98" +"76" +"54" +"32" +"10",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		},



		{
			Src:         "0x"+"fe",
			Expected: []byte{0xfe},
		},
		{
			Src:         "0x"+"fe" +"dc",
			Expected: []byte{0xfe, 0xdc},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba",
			Expected: []byte{0xfe, 0xdc, 0xba},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba" +"98",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba" +"98" +"76",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba" +"98" +"76" +"54",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba" +"98" +"76" +"54" +"32",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32},
		},
		{
			Src:         "0x"+"fe" +"dc" +"ba" +"98" +"76" +"54" +"32" +"10",
			Expected: []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		},



		{
			Src:              "01",
			Expected: []byte{0x01},
		},
		{
			Src:              "01" +"23",
			Expected: []byte{0x01, 0x23},
		},
		{
			Src:              "01" +"23" +"45",
			Expected: []byte{0x01, 0x23, 0x45},
		},
		{
			Src:              "01" +"23" +"45" +"67",
			Expected: []byte{0x01, 0x23, 0x45, 0x67},
		},
		{
			Src:              "01" +"23" +"45" +"67" +"89",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89},
		},
		{
			Src:              "01" +"23" +"45" +"67" +"89" +"AB",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB},
		},
		{
			Src:              "01" +"23" +"45" +"67" +"89" +"AB" +"CD",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD},
		},
		{
			Src:              "01" +"23" +"45" +"67" +"89" +"AB" +"CD" +"EF",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		},



		{
			Src:         "0x"+"01",
			Expected: []byte{0x01},
		},
		{
			Src:         "0x"+"01" +"23",
			Expected: []byte{0x01, 0x23},
		},
		{
			Src:         "0x"+"01" +"23" +"45",
			Expected: []byte{0x01, 0x23, 0x45},
		},
		{
			Src:         "0x"+"01" +"23" +"45" +"67",
			Expected: []byte{0x01, 0x23, 0x45, 0x67},
		},
		{
			Src:         "0x"+"01" +"23" +"45" +"67" +"89",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89},
		},
		{
			Src:         "0x"+"01" +"23" +"45" +"67" +"89" +"AB",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB},
		},
		{
			Src:         "0x"+"01" +"23" +"45" +"67" +"89" +"AB" +"CD",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD},
		},
		{
			Src:         "0x"+"01" +"23" +"45" +"67" +"89" +"AB" +"CD" +"EF",
			Expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.Parse(test.Src, "hexadecimal")
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

func TestKernel_Parse_hexadecimal_failure(t *testing.T) {

	tests := []struct{
		Src string
	}{
		{
			Src: "0x",
		},
		{
			Src: "0x0x",
		},
		{
			Src: "0x0x0x",
		},



		{
			Src: "wxyz",
		},



		{
			Src: "apple",
		},
		{
			Src: "BANANA",
		},
		{
			Src: "Cherry",
		},



		{
			Src: "MN",
		},
		{
			Src: "mn",
		},
		{
			Src: "0xMN",
		},
		{
			Src: "0xmn",
		},



		{
			Src: " 00",
		},
		{
			Src: " 12",
		},
		{
			Src: " fe",
		},



		{
			Src: "  00",
		},
		{
			Src: "  12",
		},
		{
			Src: "  fe",
		},



		{
			Src: "00  ",
		},
		{
			Src: "12  ",
		},
		{
			Src: "fe  ",
		},



		{
			Src: "00 ",
		},
		{
			Src: "12 ",
		},
		{
			Src: "fe ",
		},



		{
			Src: " 0x00",
		},
		{
			Src: " 0x12",
		},
		{
			Src: " 0xfe",
		},



		{
			Src: "  0x00",
		},
		{
			Src: "  0x12",
		},
		{
			Src: "  0xfe",
		},



		{
			Src: "0x00 ",
		},
		{
			Src: "0x12 ",
		},
		{
			Src: "0xfe ",
		},



		{
			Src: "0x00  ",
		},
		{
			Src: "0x12  ",
		},
		{
			Src: "0xfe  ",
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		err := kernel.Parse(test.Src, "hexadecimal")

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("SOURCE: %q", test.Src)
			t.Logf("ERROR: %v", err)
			continue
		}
	}
}
