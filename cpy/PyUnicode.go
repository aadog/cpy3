package cpy

import (
	"github.com/aadog/cpy3/cpy/dllimports"
	"unsafe"
)

func PyUnicode_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyUnicode_Type))
}
func PyUnicode_DecodeFSDefault(s uintptr) uintptr {
	r := defSyscallN(dllimports.PyUnicode_DecodeFSDefault, s)
	return r
}
func PyUnicode_GetLength(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyUnicode_GetLength, obj)
	return int64(r)
}
func PyUnicode_FromString(str string) uintptr {
	if str == "" {
		b := make([]byte, 1)
		b[0] = 0x0
		r := defSyscallN(dllimports.PyUnicode_FromString, uintptr(unsafe.Pointer(&b[0])))
		return r
	}
	r := defSyscallN(dllimports.PyUnicode_FromString, PascalStr(str))
	return r
}
func PyUnicode_AsUTF8(obj uintptr) string {
	r := defSyscallN(dllimports.PyUnicode_AsUTF8, obj)
	return GoStr(r)
}
