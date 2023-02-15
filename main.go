package main

import (
	"fmt"
	_ "github.com/aadog/cpy3/alib"
	"github.com/aadog/cpy3/cpy"
)

func main() {

	cpy.Py_Initialize()
	cpy.Py_SetProgramName("xxxx")
	cpy.Py_SetPath("aaa")
	cpy.Py_SetPythonHome("bb")

	fmt.Println(cpy.Py_None())
	//cpy.Py_FatalError("aaa")
	//fmt.Println(cpy.Py_CompileString("print('xx')", "", 0))
	//fp := cpy.Py_fopen_obj(cpy.Py_BuildValue("s", "test.py"), "r+")
	//cpy.PyRun_SimpleFile(fp, "text.py")

	cpy.Py_Finalize()
}
