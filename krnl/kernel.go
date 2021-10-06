package krnl

// Kernel represents the core of an address for content-addressing,
// such as with a content-addressable-storage.
//
// Example
//
// For example, imagine we had a document with just the following content:
//
//	Hello world!
//
// (We can do content-addressing in different ways.
// One way to do content-addressing is using cryptographic hash functions.
// It isn't the only way to do content-addressing.
// But it is the way we will use in this example.)
//
// (There are many different cryptographic hash functions.
// For the sake of this example, we will have to pick a cryptographic hash function to use.
// In this example we will pick the SHA3-512 cryptographic hash function.)
//
// And also, for example, we will use the SHA3-512 cryptographic hash function
// for content-addressing — then the digest, from the SHA3-512 cryptographic
// hash function, of the document with the content "Hello world!" would be:
//
//	[]byte{
//		0x95, 0xde, 0xcc, 0x72, 0xf0, 0xa5, 0x0a, 0xe4,
//		0xd9, 0xd5, 0x37, 0x8e, 0x1b, 0x22, 0x52, 0x58,
//		0x7c, 0xfc, 0x71, 0x97, 0x7e, 0x43, 0x29, 0x2c,
//		0x8f, 0x1b, 0x84, 0x64, 0x82, 0x48, 0x50, 0x9f,
//		0x1b, 0xc1, 0x8b, 0xc6, 0xf0, 0xb0, 0xd0, 0xb8,
//		0x60, 0x6a, 0x64, 0x3e, 0xff, 0x61, 0xd6, 0x11,
//		0xae, 0x84, 0xe6, 0xfb, 0xd4, 0xa2, 0x68, 0x31,
//		0x65, 0x70, 0x6b, 0xd6, 0xfd, 0x48, 0xb3, 0x34,
//	}
//
// This digest is the kernel, in this example.
//
// (We call this a “kernel” rather than a “digest” as technically, content-addressing could
// be done with a different technique, other than a cryptographic hash function. And although
// a cryptographic hash function produces a digest, other techniques might not.)
//
// One thing to note is that the kernel does not make note that the SHA3-512 was used to generate this digest.
// I.e., with just the kernel by itself, you probably have no idea what type of data is inside of it.
//
// To store the kernel with the a note about, for example, which cryptographic hash function
// that was used to generate it, use addr.Address. (addr.Address contains a krnl.Kernel)
//
// Another thing to note is that, when the kernel represents a number, the kernel does not make note of the endianness that the bytes are stored.
// So, it does not specify whether they in little-endian or big-endian.
type Kernel struct {
	loaded  bool
	storage [1024]byte
	length  int
}

// krnl.Kernel is an option type. Thus is can be ‘nothing’, or ‘something’.
//
// This is useful, for example, to check if the receiver of a method is or isn't ‘nothing’.
//
// For example:
//
//	if nothing() == receiver {
//		//@TODO
//	}
//
func nothing() Kernel {
	return Kernel{}
}
