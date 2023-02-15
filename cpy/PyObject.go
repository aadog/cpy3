package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyObject_GetAttrString(obj uintptr, attr_name string) uintptr {
	r := defSyscallN(dllimports.PyObject_GetAttrString, obj, PascalStr(attr_name))
	return r
}
func PyObject_HasAttrString(obj uintptr, attr_name string) int {
	r := defSyscallN(dllimports.PyObject_HasAttr, obj, PascalStr(attr_name))
	return int(r)
}
func PyObject_SetAttrString(obj uintptr, attr_name string, v uintptr) int {
	r := defSyscallN(dllimports.PyObject_SetAttr, obj, PascalStr(attr_name), v)
	return int(r)
}

func PyObject_DelAttrString(obj uintptr, attr_name string) int {
	r := defSyscallN(dllimports.PyObject_DelAttrString, obj, PascalStr(attr_name))
	return int(r)
}

func PyObject_Type(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_Type, obj)
	return r
}

func PyObject_Str(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_Str, obj)
	return r
}

func PyObject_GetAttr(obj uintptr, attr_name uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_GetAttr, obj, attr_name)
	return r
}

func PyObject_Call(obj uintptr, args uintptr, kwargs uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_Call, obj, args, kwargs)
	return r
}
func PyObject_CallObject(obj uintptr, args uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_CallObject, obj, args)
	return r
}
func PyObject_CallNoArgs(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_CallNoArgs, obj)
	return r
}
