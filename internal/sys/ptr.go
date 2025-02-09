package sys

import (
	"unsafe"

	"github.com/kubeshark/ebpf/internal/unix"
)

// NewPointer creates a 64-bit pointer from an unsafe Pointer.
func NewPointer(ptr unsafe.Pointer) Pointer {
	return Pointer{ptr: ptr}
}

// NewSlicePointer creates a 64-bit pointer from a byte slice.
func NewSlicePointer(buf []byte) Pointer {
	if len(buf) == 0 {
		return Pointer{}
	}

	return Pointer{ptr: unsafe.Pointer(&buf[0])}
}

// NewSlicePointer creates a 64-bit pointer from a byte slice.
//
// Useful to assign both the pointer and the length in one go.
func NewSlicePointerLen(buf []byte) (Pointer, uint32) {
	return NewSlicePointer(buf), uint32(len(buf))
}

// NewStringPointer creates a 64-bit pointer from a string.
func NewStringPointer(str string) Pointer {
	p, err := unix.BytePtrFromString(str)
	if err != nil {
		return Pointer{}
	}

	return Pointer{ptr: unsafe.Pointer(p)}
}
