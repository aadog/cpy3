package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

var py_NoneStruct = dllimports.Py_NoneStruct

func Py_None() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.Py_NoneStruct))
}
