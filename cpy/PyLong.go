package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

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

func PyLong_FromLong(n int) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLong, uintptr(n))
	return r
}

func PyLong_FromLongLong(n int64) uintptr {
	r := defSyscallN(dllimports.PyLong_FromLongLong, uintptr(n))
	return r
}
func PyLong_FromDouble(n float64) uintptr {
	r := defSyscallN(dllimports.PyLong_FromDouble, uintptr(n))
	return r
}

//func PyLong_FromString(n float64) uintptr {
//	r, _, _ := pyLong_FromLong.Call(uintptr(n))
//	return r
//}
