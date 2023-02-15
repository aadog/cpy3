package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyBytes_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyBytes_Type))
}

func PyBytes_FromString(s string) uintptr {
	r := defSyscallN(dllimports.PyBytes_AsString, PascalStr(s))
	return r
}
