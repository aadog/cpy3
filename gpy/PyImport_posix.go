package gpy

import "github.com/aadog/cpy3/cpy"

func PyImport_Import(name string) *PyModule {
	pName := NewPyUnicode(name)
	return AsPyModule(cpy.PyImport_Import(pName._instance()))
}
func PyImport_AppendInittab(name string, initFunc uintptr) int {
	r := cpy.PyImport_AppendInittab(name, initFunc)
	return r
}
