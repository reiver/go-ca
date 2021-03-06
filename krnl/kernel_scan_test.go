package krnl_test

import (
	"github.com/reiver/go-ca/krnl"

	"bytes"
	"time"

	"testing"
)

func TestKernel_Scan_success(t *testing.T) {

	tests := []struct{
		Src interface{}
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
			Src:      "",
			Expected: []byte{},
		},
		{
			Src:      krnl.Kernel{},
			Expected: []byte{},
		},
		{
			Src:      new(krnl.Kernel),
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
			Src:      []byte{0xf9},
			Expected: []byte{0xf9},
		},
		{
			Src:      []byte{0xfa},
			Expected: []byte{0xfa},
		},
		{
			Src:      []byte{0xfb},
			Expected: []byte{0xfb},
		},
		{
			Src:      []byte{0xfc},
			Expected: []byte{0xfc},
		},
		{
			Src:      []byte{0xfd},
			Expected: []byte{0xfd},
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
			Src:            "\x00",
			Expected: []byte{0x00},
		},
		{
			Src:            "\x01",
			Expected: []byte{0x01},
		},
		{
			Src:            "\x02",
			Expected: []byte{0x02},
		},
		{
			Src:            "\x03",
			Expected: []byte{0x03},
		},
		{
			Src:            "\x04",
			Expected: []byte{0x04},
		},
		{
			Src:            "\x05",
			Expected: []byte{0x05},
		},

		{
			Src:            "\x12",
			Expected: []byte{0x12},
		},

		{
			Src:            "\x2a",
			Expected: []byte{0x2a},
		},

		{
			Src:            "\xf9",
			Expected: []byte{0xf9},
		},
		{
			Src:            "\xfa",
			Expected: []byte{0xfa},
		},
		{
			Src:            "\xfb",
			Expected: []byte{0xfb},
		},
		{
			Src:            "\xfc",
			Expected: []byte{0xfc},
		},
		{
			Src:            "\xfd",
			Expected: []byte{0xfd},
		},
		{
			Src:            "\xfe",
			Expected: []byte{0xfe},
		},
		{
			Src:            "\xff",
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
		{
			Src:
				"\x95\xde\xcc\x72\xf0\xa5\x0a\xe4"+
				"\xd9\xd5\x37\x8e\x1b\x22\x52\x58"+
				"\x7c\xfc\x71\x97\x7e\x43\x29\x2c"+
				"\x8f\x1b\x84\x64\x82\x48\x50\x9f"+
				"\x1b\xc1\x8b\xc6\xf0\xb0\xd0\xb8"+
				"\x60\x6a\x64\x3e\xff\x61\xd6\x11"+
				"\xae\x84\xe6\xfb\xd4\xa2\x68\x31"+
				"\x65\x70\x6b\xd6\xfd\x48\xb3\x34",
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
		{
			Src: func() krnl.Kernel {

				var kernel krnl.Kernel

				err := kernel.Scan([]byte{
					0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
					0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
					0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
					0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
					0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
					0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
					0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
					0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
				})
				if nil != err {
					panic(err)
				}

				return kernel
			}(),
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
		{
			Src: func() *krnl.Kernel {

				var kernel krnl.Kernel

				err := kernel.Scan([]byte{
					0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
					0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
					0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
					0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
					0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
					0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
					0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
					0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
				})
				if nil != err {
					panic(err)
				}

				return &kernel
			}(),
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
			err := kernel.Scan(test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		{
			var expected []byte = test.Expected
			var actual   []byte = kernel.Bytes()

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual value from .Bytes() method is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}
	}
}

func TestKernel_Scan_failure(t *testing.T) {

	tests := []struct{
		Src interface{}
	}{
		{
			Src: int64(-12345),
		},
		{
			Src: int64(-1),
		},
		{
			Src: int64(0),
		},
		{
			Src: int64(1),
		},
		{
			Src: int64(12345),
		},



		{
			Src: -123.45,
		},
		{
			Src: -1.0,
		},
		{
			Src: 0.0,
		},
		{
			Src: 1.0,
		},
		{
			Src: 123.45,
		},



		{
			Src: false,
		},
		{
			Src: true,
		},



		{
			Src: time.Now(),
		},



		{
			Src: nil,
		},



		{
			Src: struct{
				X int
				Y int
			}{
				X:  5,
				Y: -1,
			},
		},
	}

	for testNumber, test := range tests {

		var kernel krnl.Kernel

		err := kernel.Scan(test.Src)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR: %#v", err)
			continue
		}
	}
}
