package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyType_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyType_Type))
}

func PyType_Check(obj uintptr) bool {
	r := defSyscallN(dllimports.PyType_Check, obj)
	return GoBool(r)
}

func PyType_IsSubtype(a uintptr, b uintptr) bool {
	r := defSyscallN(dllimports.PyType_IsSubtype, a, b)
	return GoBool(r)
}

func PyType_GetName(tp uintptr) string {
	r := defSyscallN(dllimports.PyType_GetName, tp)
	defer Py_DecRef(r)
	return GoStr(r)
}
