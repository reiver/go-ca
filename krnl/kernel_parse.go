package krnl

// Parse loads Kernel with the value parsed from ‘value’ using the parsing algoritm specified by ‘format’.
//
// For example:
//
//	var kernel krnl.Kernel
//	
//	// ...
//	
//	var value string = "ld7McvClCuTZ1TeOGyJSWHz8cZd-QyksjxuEZIJIUJ8bwYvG8LDQuGBqZD7_YdYRroTm-9SiaDFlcGvW_UizNA"
//	var format string = "base64url"
//	
//	// ...
//	
//	err := kernel.Parse(value, format)
//
// Some of the formats that Parse supports is:
//
// • asci85
//
// • base32
//
// • base64
//
// • base64url
//
// • hexadecimal
func (receiver *Kernel) Parse(value string, format string) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch format {
	case "ascii85":
		return receiver.parse_ascii85(value)
	case "base32":
		return receiver.parse_base32(value)
	case "base64":
		return receiver.parse_base64(value)
	case "base64url":
		return receiver.parse_base64url(value)
	case "hexadecimal":
		return receiver.parse_hexadecimal(value)

	default:
		return errorf("unsupported format — %q", format)
	}
}
