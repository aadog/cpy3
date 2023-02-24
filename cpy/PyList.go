package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyList_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyList_Type))
}

// Return value: New reference.
func PyList_New(len int64) uintptr {
	r := defSyscallN(dllimports.PyList_New, uintptr(len))
	return r
}

func PyList_SetItem(obj uintptr, index int64, item uintptr) int {
	r := defSyscallN(dllimports.PyList_SetItem, obj, uintptr(index), item)
	return int(r)
}

// Return value: Borrowed reference.
func PyList_GetItem(obj uintptr, index int64) uintptr {
	r := defSyscallN(dllimports.PyList_GetItem, obj, uintptr(index))
	return r
}

func PyList_GetSlice(obj uintptr, low int64, high int64) uintptr {
	r := defSyscallN(dllimports.PyList_GetSlice, obj, uintptr(low), uintptr(high))
	return r
}
func PyList_Size(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyList_Size, obj)
	return int64(r)
}
func PyList_Insert(obj uintptr, index int64, item uintptr) int {
	r := defSyscallN(dllimports.PyList_Insert, obj, uintptr(index), item)
	return int(r)
}
func PyList_Append(obj uintptr, item uintptr) int {
	r := defSyscallN(dllimports.PyList_Append, obj, item)
	return int(r)
}
