package cpy

import (
	"github.com/aadog/cpy3/cpy/dllimports"
	"unsafe"
)

func PyModule_Type() uintptr {
	return uintptr(dllimports.GetImportDefFunc(uiLib, dllimports.PyModule_Type))
}

type PyMethodDef struct {
	Ml_name  uintptr
	Ml_meth  uintptr
	Ml_flags int
	Ml_doc   uintptr
}
type PyObject struct {
	//Ob_next   uintptr
	//Ob_prev   uintptr
	Ob_refcnt uintptr
	Ob_type   uintptr
}

type PyModuleDef_Base struct {
	Ob_base PyObject
	M_init  uintptr
	M_index uintptr
	M_copy  uintptr
}
type PyModuleDef struct {
	M_base     PyModuleDef_Base
	M_name     uintptr
	M_doc      uintptr
	M_size     int64
	M_methods  uintptr
	M_slots    uintptr
	M_traverse uintptr
	M_clear    uintptr
	M_free     uintptr
}

func PyObjectFromPtr(ptr uintptr) *PyObject {
	return (*PyObject)(unsafe.Pointer(ptr))
}

func PyObject_HEAD_INIT(ob_type uintptr) PyObject {
	return PyObject{
		Ob_refcnt: 1,
		Ob_type:   0,
	}
}

func PyModule_Create2(PyModuleDef uintptr, apiver int) uintptr {
	r := defSyscallN(dllimports.PyModule_Create2, PyModuleDef, uintptr(apiver))
	return r
}
func PyImport_Import(name uintptr) uintptr {
	r := defSyscallN(dllimports.PyImport_Import, name)
	return r
}

func PyModule_GetName(obj uintptr) string {
	r := defSyscallN(dllimports.PyModule_GetName, obj)
	return GoStr(r)
}
func PyModule_GetDict(obj uintptr) uintptr {
	r := defSyscallN(dllimports.PyModule_GetDict, obj)
	return r
}
func PyModule_AddFunctions(obj uintptr, functionsDef uintptr) int {
	r := defSyscallN(dllimports.PyModule_AddFunctions, obj, functionsDef)
	return int(r)
}
func PyModule_AddIntConstant(obj uintptr, name string, value int64) int {
	r := defSyscallN(dllimports.PyModule_AddIntConstant, obj, PascalStr(name), uintptr(value))
	return int(r)
}
func PyModule_AddStringConstant(obj uintptr, name string, value string) int {
	r := defSyscallN(dllimports.PyModule_AddStringConstant, obj, PascalStr(name), PascalStr(value))
	return int(r)
}
func PyModule_AddObject(obj uintptr, name string, value uintptr) int {
	r := defSyscallN(dllimports.PyModule_AddIntConstant, obj, PascalStr(name), uintptr(value))
	return int(r)
}
func PyModule_AddObjectRef(obj uintptr, name string, value uintptr) int {
	r := defSyscallN(dllimports.PyModule_AddObjectRef, obj, PascalStr(name), uintptr(value))
	return int(r)
}
