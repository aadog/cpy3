package cpy

import (
	"github.com/aadog/cpy3/cpy/dllimports"
)

func PyLong_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyLong_Type))
}

func PyLong_AsDouble(obj uintptr) float64 {
	r := defSyscallN(dllimports.PyLong_AsDouble, obj)
	return float64(r)
}
func PyLong_AsLong(obj uintptr) int {
	r := defSyscallN(dllimports.PyLong_AsLong, obj)
	return int(r)
}
func PyLong_AsLongLong(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyLong_AsLongLong, obj)
	return int64(r)
}
func PyLong_AsUnsignedLong(obj uintptr) uint {
	r := defSyscallN(dllimports.PyLong_AsUnsignedLong, obj)
	return uint(r)
}
func PyLong_AsUnsignedLongLong(obj uintptr) uint64 {
	r := defSyscallN(dllimports.PyLong_AsUnsignedLongLong, obj)
	return uint64(r)
}

// Return value: New reference.
func PyLong_FromLong(n int) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLong, uintptr(n))
	return r
}

// Return value: New reference.
func PyLong_FromLongLong(n int64) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(n))
	return r
}

// Return value: New reference.
func PyLong_FromDouble(n float64) uintptr {
	//浮点传不过去
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(n))
	return r
}

// Return value: New reference.
func PyLong_FromFloat32(n float32) uintptr {
	//浮点传不过去
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(n))
	return r
}

// Return value: New reference.
func PyLong_FromUnicodeObject(obj uintptr,base int) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(obj),uintptr(base))
	return r
}

// Return value: New reference.
func PyLong_FromVoidPtr(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(obj))
	return r
}

//func PyLong_FromString(n float64) uintptr {
//	r, _, _ := pyLong_FromLong.Call(uintptr(n))
//	return r
//}
