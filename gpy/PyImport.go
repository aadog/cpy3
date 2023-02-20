package gpy

import "github.com/aadog/cpy3/cpy"

func PyImport_Import(name string) *PyModule {
	pName := NewPyUnicode(name)
	return AsPyModule(cpy.PyImport_Import(pName._instance()))
}
func PyImport_AppendInittab(name string, initFunc func() *PyGoModule) int {
	r := cpy.PyImport_AppendInittab(name, NewModuleInitFuncCallBack(name, func() *PyGoModule {
		obj := initFunc()
		return obj
	}))
	return r
}
