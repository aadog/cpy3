package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyPyInstanceMethod_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyInstanceMethod_Type))
}

func PyInstanceMethod_New(cfn uintptr) uintptr {
	r := defSyscallN(dllimports.PyInstanceMethod_New, cfn)
	return r
}
