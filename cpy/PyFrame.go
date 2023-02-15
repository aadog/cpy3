package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyFrame_GetBack(frame uintptr) uintptr {
	r := defSyscallN(dllimports.PyFrame_GetBack, frame)
	return r
}
func PyFrame_GetCode(frame uintptr) uintptr {
	r := defSyscallN(dllimports.PyFrame_GetCode)
	return r
}
