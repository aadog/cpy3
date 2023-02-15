package gpy

import "github.com/aadog/cpy3/cpy"

func PyImport_Import(name string) *PyModule {
	pName := PyUnicode_DecodeFSDefault(name)
	defer pName.DecRef()
	return NewPyModuleWithPtr(cpy.PyImport_Import(pName._instance()))
}

func PyImport_AppendInittab(name string, initFunc func() IPyObject) int {
	r := cpy.PyImport_AppendInittab(name, NewModuleInitFuncCallBack(name, func() IPyObject {
		obj := initFunc()
		return obj
	}))
	return r
}
