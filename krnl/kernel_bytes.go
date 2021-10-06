package krnl

// slice returns a slice of krnl.Kernel.storage of length krnl.Kernel.length
//
// In many cases, only part of the krnl.Kernel.storage array is storing actual data.
// krnl.Kernel.length tells us how much of krnl.Kernel.storage is actually being used to store data.
//
// To get at just the valid part of krnl.Kernel.storage, we use slice.
//
// This returns the actual memory in the krnl.Kernel.
// I.e., if you modify this it will modify the value of the kernel.
func (receiver *Kernel) slice() []byte {
	if nil == receiver {
		return nil
	}

	return receiver.storage[:receiver.length]
}

// Bytes returns the kernel as []byte.
//
// Note that you cannot change the bytes that are actually in the kernel by changing the bytes returned by Bytes.
//
// This is how one might use it:
//
//	var kernel krnl.Kernel
//	
//	// ...
//	
//	err := kernel.Parse("ld7McvClCuTZ1TeOGyJSWHz8cZd-QyksjxuEZIJIUJ8bwYvG8LDQuGBqZD7_YdYRroTm-9SiaDFlcGvW_UizNA", "base64url")
//	
//	// ...
//	
//	var p []byte = kernel.Bytes()
func (receiver Kernel) Bytes() []byte {
	if nothing() == receiver {
		return nil
	}

	// Note that since this does NOT have a pointer-receiever, and is a copy,
	// that returning receiver.slice() is OK, since the slice does NOT point to the
	// original memory.

	return receiver.slice()
}
