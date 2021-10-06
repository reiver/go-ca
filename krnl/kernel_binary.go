package krnl

// MarshalBinary returns the value of Kernel as a []byte.
//
// MarshalBinary also makes krnl.Kernel fit the encoding.BinaryMarshaler interface.
func (receiver Kernel) MarshalBinary() ([]byte, error) {
	if nothing() == receiver {
		return nil, errNothing
	}

	// Note that since this does NOT have a pointer-receiever, and is a copy,
	// that returning the value of receiver.slice() is OK,
	// as the slice does NOT point to the original memory.

	return receiver.slice(), nil
}

// UnmarshalBinary loads Kernel with the value of the []byte.
//
// UnmarshalBinary also makes krnl.Kernel fit the encoding.BinaryUnmarshaler interface.
//
// For example:
//
//	var kernel krnl.Kernel
//	
//	// ...
//	
//	var p []byte = ??? //@TODO
//
//	// ...
//	
//	err := kernel.UnmarshalBinary(p)
//
// UnmarshalBinary also makes krnl.Kernel fit the encoding.BinaryUnmarshaler interface.
func (receiver *Kernel) UnmarshalBinary(p []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	actualLength := len(p)
	maxLength := len(receiver.storage)

	if actualLength > maxLength {
		return errorf("source []byte too big — max_length=%d actual_length=%d", maxLength, actualLength)
	}

	receiver.loaded = true
	receiver.length = copy(receiver.storage[:], p)

	return nil
}
