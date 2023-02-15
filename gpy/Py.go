package gpy

import (
	"fmt"
	"github.com/aadog/cpy3/cpy"
	"sync"
)

var SystemModuleMap = sync.Map{}
var GoModule *PyModule

func InitGoModuleName(name string, doc string) {
	PyImport_AppendInittab(name, func() IPyObject {
		GoModule = CreateModule(name, doc)
		cls := CreateClass("MyClass", "")
		cls.AddFunction("z", func(self *PyObject) {
			fmt.Println("z")
		})
		GoModule.AddClass(cls)

		return GoModule
	})
}
func Initialize() {
	//init python3
	cpy.Py_Initialize()

	//new userexception
	_UserException = PyErr_NewException("gofunction.error", PyNil, PyNil)
}

func IsInitialized() int {
	return cpy.Py_IsInitialized()
}
func Finalize() {
	SystemModuleMap.Range(func(key, value any) bool {
		//SystemModuleMap.Delete(key)
		m := value.(*PyModule)
		m.DecRef()
		return true
	})
	_UserException.DecRef()
	cpy.Py_Finalize()
}
func FinalizeEx() int {
	SystemModuleMap.Range(func(key, value any) bool {
		//SystemModuleMap.Delete(key)
		m := value.(*PyModule)
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
