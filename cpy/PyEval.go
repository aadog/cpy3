package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyEval_GetBuiltins() uintptr {
	r := defSyscallN(dllimports.PyEval_GetBuiltins)
	return r
}
func PyEval_GetLocals() uintptr {
	r := defSyscallN(dllimports.PyEval_GetLocals)
	return r
}

func PyEval_GetGlobals() uintptr {
	r := defSyscallN(dllimports.PyEval_GetGlobals)
	return r
}
func PyEval_GetFrame() uintptr {
	r := defSyscallN(dllimports.PyEval_GetFrame)
	return r
}
