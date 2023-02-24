package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyErr_NewException(name string, base uintptr, dict uintptr) uintptr {
	r := defSyscallN(dllimports.PyErr_NewException, PascalStr(name), base, dict)
	return r
}
func PyExc_Exception() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyExc_Exception))
}

func PyExc_ValueError() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyExc_ValueError))
}

func PyErr_SetString(tp uintptr, message string) {
	defSyscallN(dllimports.PyErr_SetString, tp, PascalStr(message))
}

//Return value: Borrowed reference.
func PyErr_Occurred()uintptr{
	return defSyscallN(dllimports.PyErr_Occurred)
}

