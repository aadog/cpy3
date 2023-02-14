package api

import (
	"github.com/aadog/cpy3/cpy/api/dllimports"
	"unsafe"
)

func PyUnicode_DecodeFSDefault(s uintptr) uintptr {
	r := defSyscallN(dllimports.PyUnicode_DecodeFSDefault, s)
	return r
}
func PyUnicode_GetLength(obj uintptr) int64 {
	r := defSyscallN(dllimports.PyUnicode_GetLength, obj)
	return int64(r)
}
func PyUnicode_FromString(u string) uintptr {
	if u == "" {
		b := make([]byte, 1)
		b[0] = 0x0
		r := defSyscallN(dllimports.PyUnicode_FromString, uintptr(unsafe.Pointer(&b[0])))
		return r
	}
	r := defSyscallN(dllimports.PyUnicode_FromString, PascalStr(u))
	return r
}
func PyUnicode_AsUTF8(obj uintptr) string {
	r := defSyscallN(dllimports.PyUnicode_AsUTF8, obj)
	return GoStr(r)
}
