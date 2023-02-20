package gpy

import (
	"unsafe"
)

// IPyObject 共公的对象接口
type IPyObject interface {
	DecRef()
	IncRef()
	RefCount() int
	Instance() uintptr
	IsValid() bool
	UnsafeAddr() unsafe.Pointer

	ClassName() string
	Free()
	GetHashCode() int64
	Equals(IPyObject) bool
	//DisposeOf()
	Type() *PyType
	InstanceSize() int32
	//InheritsFrom(types.TClass) bool

	//Is() TIs
	//As() TAs

	ToString() string
	Str() string
}
