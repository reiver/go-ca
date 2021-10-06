package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"
	"math/rand"

	"testing"
)

func TestKernel_MarshalBinary_success(t *testing.T) {

	tests := []struct{
		Src []byte
		Expected []byte
	}{
		{
			Src:      []byte(nil),
			Expected: []byte{},
		},
		{
			Src:      []byte{},
			Expected: []byte{},
		},



		{
			Src:      []byte{0x00},
			Expected: []byte{0x00},
		},
		{
			Src:      []byte{0x01},
			Expected: []byte{0x01},
		},
		{
			Src:      []byte{0x02},
			Expected: []byte{0x02},
		},
		{
			Src:      []byte{0x03},
			Expected: []byte{0x03},
		},
		{
			Src:      []byte{0x04},
			Expected: []byte{0x04},
		},
		{
			Src:      []byte{0x05},
			Expected: []byte{0x05},
		},

		{
			Src:      []byte{0x12},
			Expected: []byte{0x12},
		},

		{
			Src:      []byte{0x2a},
			Expected: []byte{0x2a},
		},

		{
			Src:      []byte{0xfe},
			Expected: []byte{0xfe},
		},
		{
			Src:      []byte{0xff},
			Expected: []byte{0xff},
		},



		{
			Src:      []byte{0x00,0x00},
			Expected: []byte{0x00,0x00},
		},
		{
			Src:      []byte{0x01,0x00},
			Expected: []byte{0x01,0x00},
		},
		{
			Src:      []byte{0x02,0x00},
			Expected: []byte{0x02,0x00},
		},
		{
			Src:      []byte{0x03,0x00},
			Expected: []byte{0x03,0x00},
		},
		{
			Src:      []byte{0x04,0x00},
			Expected: []byte{0x04,0x00},
		},
		{
			Src:      []byte{0x05,0x00},
			Expected: []byte{0x05,0x00},
		},

		{
			Src:      []byte{0x00,0x01},
			Expected: []byte{0x00,0x01},
		},
		{
			Src:      []byte{0x01,0x01},
			Expected: []byte{0x01,0x01},
		},
		{
			Src:      []byte{0x02,0x01},
			Expected: []byte{0x02,0x01},
		},
		{
			Src:      []byte{0x03,0x01},
			Expected: []byte{0x03,0x01},
		},
		{
			Src:      []byte{0x04,0x01},
			Expected: []byte{0x04,0x01},
		},
		{
			Src:      []byte{0x05,0x01},
			Expected: []byte{0x05,0x01},
		},

		{
			Src:      []byte{0x12,0x34},
			Expected: []byte{0x12,0x34},
		},



		{
			Src:      []byte{0xfa,0xfe},
			Expected: []byte{0xfa,0xfe},
		},
		{
			Src:      []byte{0xfb,0xfe},
			Expected: []byte{0xfb,0xfe},
		},
		{
			Src:      []byte{0xfc,0xfe},
			Expected: []byte{0xfc,0xfe},
		},
		{
			Src:      []byte{0xfd,0xfe},
			Expected: []byte{0xfd,0xfe},
		},
		{
			Src:      []byte{0xfe,0xfe},
			Expected: []byte{0xfe,0xfe},
		},
		{
			Src:      []byte{0xff,0xfe},
			Expected: []byte{0xff,0xfe},
		},



		{
			Src:      []byte{0xfa,0xff},
			Expected: []byte{0xfa,0xff},
		},
		{
			Src:      []byte{0xfb,0xff},
			Expected: []byte{0xfb,0xff},
		},
		{
			Src:      []byte{0xfc,0xff},
			Expected: []byte{0xfc,0xff},
		},
		{
			Src:      []byte{0xfd,0xff},
			Expected: []byte{0xfd,0xff},
		},
		{
			Src:      []byte{0xfe,0xff},
			Expected: []byte{0xfe,0xff},
		},
		{
			Src:      []byte{0xff,0xff},
			Expected: []byte{0xff,0xff},
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
			Expected: []byte{
				0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
				0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
				0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
				0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
				0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
				0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
				0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
				0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
			},
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		{
			err := kernel.UnmarshalBinary(test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error when unmarshaling-binary but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		var actual []byte
		{
			var err error

			actual, err = kernel.MarshalBinary()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error when marshaling-binary but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
			if nil == actual {
				t.Errorf("For test #%d, did not expect marshaled-binary []byte to be nil but actually got was.", testNumber)
				continue
			}
		}

		{
			var expected []byte = test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual bytes from marshaling-binary is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}
	}
}

func TestKernel_MarshalBinary_success1024(t *testing.T) {

	const max = 1024

	for testNumber:=0; testNumber<max; testNumber++ {

		var src []byte = make([]byte, testNumber)
		{
			if nil == src {
				t.Errorf("For test #%d, did not expect the made []byte to be nil but actually was.", testNumber, )
				continue
			}

			for i:=0; i<len(src); i++ {
				src[i] = byte(rand.Int() % 256)
			}

		}

		var kernel krnl.Kernel
		{
			err := kernel.UnmarshalBinary(src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		var actual []byte
		{
			var err error

			actual, err = kernel.MarshalBinary()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
			if nil == actual {
				t.Errorf("For test #%d, did not expected marshaled-binary to be nil but actually was.", testNumber)
				t.Logf("BYTES: %#v", actual)
				continue
			}
		}

		if expected := src; !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled-binary is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("EXPECTED LENGTH: %d", len(expected))
			t.Logf("ACTUAL   LENGTH: %d", len(actual))
			continue
		}
	}
}

func TestKernel_MarshalBinary_failureNothing(t *testing.T) {

	var kernel krnl.Kernel

	var err error
	{
		var p []byte

		p, err = kernel.MarshalBinary()
		if nil != p {
			t.Errorf("Expected marshaled-binary []byte to be nil but actually wasn't.")
			return
		}

	}

	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		t.Logf("ERROR: %#v", err)
		return
	}
}
