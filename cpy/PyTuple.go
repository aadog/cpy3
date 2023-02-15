package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyTuple_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyTuple_Type))
}

func PyTuple_Size(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyTuple_Size, obj)
	return int64(r)
}
func PyTuple_Check(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyTuple_Check, obj)
	return int64(r)
}

func PyTuple_GetItem(obj uintptr, pos int64) uintptr {
	r := defSyscallN(dllimports.PyTuple_GetItem, obj, uintptr(pos))
	return r
}
func PyTuple_SetItem(obj uintptr, pos int64, o uintptr) uintptr {
	r := defSyscallN(dllimports.PyTuple_SetItem, obj, uintptr(pos), o)
	return r
}

func PyTuple_GetSlice(obj uintptr, low int64, high int64) uintptr {
	r := defSyscallN(dllimports.PyTuple_GetSlice, obj, uintptr(low), uintptr(high))
	return r
}

func PyTuple_New(len int64) uintptr {
	r := defSyscallN(dllimports.PyTuple_New, uintptr(len))
	return r
}
