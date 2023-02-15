package gpy

import "github.com/aadog/cpy3/cpy"

var _UserException *PyObject

func UserException() *PyObject {
	return _UserException
}

func PyErr_NewException(name string, base *PyObject, dict *PyObject) *PyObject {
	return NewPyObjectWithPtr(cpy.PyErr_NewException(name, base._instance(), dict._instance()))
}

func PyExc_Exception() *PyObject {
	return AsPyObject(cpy.PyExc_Exception())
}
func PyExc_ValueError() *PyObject {
	return AsPyObject(cpy.PyExc_ValueError())
}

func PyErr_SetString(tp *PyObject, message string) {
	cpy.PyErr_SetString(tp._instance(), message)
}
