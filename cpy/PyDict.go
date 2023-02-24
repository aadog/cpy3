package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyDict_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyDict_Type))
}

func PyDict_SetItem(obj uintptr, key uintptr, val uintptr) int {
	r := defSyscallN(dllimports.PyDict_SetItem, obj, key, val)
	return int(r)
}

func PyDict_SetItemString(obj uintptr, key string, val uintptr) int {
	r := defSyscallN(dllimports.PyDict_SetItemString, obj, PascalStr(key), val)
	return int(r)
}
// Return value: New reference.
func PyDict_New() uintptr {
	r := defSyscallN(dllimports.PyDict_New)
	return r
}

func PyDict_Size(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyDict_Size, obj)
	return int64(r)
}

func PyDict_Clear(obj uintptr) {
	defSyscallN(dllimports.PyDict_Clear, obj)
}

// Return value: Borrowed reference.
func PyDict_GetItem(obj uintptr, key uintptr) uintptr {
	r := defSyscallN(dllimports.PyDict_GetItem, obj, key)
	return r
}
// Return value: New reference.
func PyDict_Keys(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyDict_Keys, obj)
	return r
}
// Return value: Borrowed reference.
func PyDict_GetItemString(obj uintptr, key string) uintptr {
	r := defSyscallN(dllimports.PyDict_GetItemString, obj, PascalStr(key))
	return r
}

func PyDict_DelItem(obj uintptr, key uintptr) int {
	r := defSyscallN(dllimports.PyDict_DelItem, obj, key)
	return int(r)
}
func PyDict_DelItemString(obj uintptr, key string) int {
	r := defSyscallN(dllimports.PyDict_DelItemString, obj, PascalStr(key))
	return int(r)
}
