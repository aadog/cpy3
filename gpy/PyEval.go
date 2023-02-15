package gpy

import "github.com/aadog/cpy3/cpy"

func PyEval_GetBuiltins() *PyObject {
	return AsPyObject(cpy.PyEval_GetBuiltins())
}
func PyEval_GetLocals() *PyObject {
	return AsPyObject(cpy.PyEval_GetLocals())
}
func PyEval_GetGlobals() *PyObject {
	return AsPyObject(cpy.PyEval_GetGlobals())
}
func PyEval_GetFrame() *PyFrame {
	return AsPyFrame(cpy.PyEval_GetFrame())
}
