package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyBool_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyBool_Type))
}

func PyBool_FromLong(b bool) uintptr {
	r := defSyscallN(dllimports.PyBool_Check, PascalBool(b))
	return r
}

func PyBool_Check(obj uintptr) bool {
	r := defSyscallN(dllimports.PyBool_Check, obj)
	return GoBool(r)
}
