package gpy

import "unsafe"

var (
	// nullptr
	nullptr unsafe.Pointer = nil //unsafe.Pointer(uintptr(0))
	PyNil                  = AsPyNil()
)
