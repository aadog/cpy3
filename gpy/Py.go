package gpy

import (
	//"fmt"
	"github.com/aadog/cpy3/cpy"
	"sync"
)

var SystemModuleMap = sync.Map{}

var GoModule *PyModule

func Initialize() {
	//init python3
	cpy.Py_Initialize()

	//new userexception
	_UserException = PyErr_NewException("go.error", PyNil, PyNil)
}

func IsInitialized() int {
	return cpy.Py_IsInitialized()
}
func Finalize() {
	SystemModuleMap.Range(func(key, value any) bool {
		SystemModuleMap.Delete(key)
		m := value.(*PyGoModule)
		m.DecRef()
		return true
	})
	_UserException.Free()
	cpy.Py_Finalize()
}
func FinalizeEx() int {
	SystemModuleMap.Range(func(key, value any) bool {
		SystemModuleMap.Delete(key)
		m := value.(*PyGoModule)
		m.DecRef()
		return true
	})
	_UserException.DecRef()
	return cpy.Py_FinalizeEx()
}
func SetProgramName(name string) {
	cpy.Py_SetProgramName(name)
}
func SetPythonHome(home string) {
	cpy.Py_SetPythonHome(home)
}
func SetPath(name string) {
	cpy.Py_SetPath(name)
}
