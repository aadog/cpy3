package cpy

import (
	"github.com/aadog/cpy3/cpy/dllimports"
)

func PyFloat_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyFloat_Type))
}

func PyFloat_Check(obj uintptr) int {
	r := defSyscallN(dllimports.PyFloat_Check, obj)
	return int(r)
}
// Return value: New reference
func PyFloat_FromString(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyFloat_FromString, obj)
	return r
}
func PyFloat_FromDouble(v float64) uintptr {
	r := defSyscallN(dllimports.PyFloat_FromDouble, uintptr(v))
	return r
}
func PyFloat_AsDouble(obj uintptr) float64 {
	r := defSyscallN(dllimports.PyFloat_AsDouble, uintptr(obj))
	return float64(r)
}
