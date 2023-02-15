package cpy

import "github.com/aadog/cpy3/cpy/dllimports"

func PyObject_New(tp uintptr) uintptr {
	return defSyscallN(dllimports.PyObject_New, tp)
}

// Return value: New reference
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

// Return value: New reference.
func PyObject_Type(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_Type, obj)
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

func PyObject_Hash(obj uintptr) int {
	r := defSyscallN(dllimports.PyObject_Hash, obj)
	return int(r)
}

func PyObject_Str(obj uintptr) string {
	r := defSyscallN(dllimports.PyObject_Str, obj)
	defer Py_DecRef(r)
	return PyUnicode_AsUTF8(r)
}

// Return value: New reference.
func PyObject_Dir(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyObject_Dir, obj)
	return r
}

func PyObject_Del(obj uintptr) {
	defSyscallN(dllimports.PyObject_Del, obj)
}
