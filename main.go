package main

import (
	"fmt"
	_ "github.com/aadog/cpy3/alib"
	"github.com/aadog/cpy3/cpy"
	"github.com/aadog/cpy3/gpy"
	"runtime"
)

func main() {

	cpy.Py_Initialize()
	cpy.Py_SetProgramName("xxxx")
	cpy.Py_SetPath("aaa")
	cpy.Py_SetPythonHome("bb")

	fmt.Println(gpy.PyLong_FromDouble(100.011).AsInt64())
	runtime.GC()
	//time.Sleep(time.Second * 1000)
	//cpy.Py_FatalError("aaa")
	//fmt.Println(cpy.Py_CompileString("print('xx')", "", 0))
	//fp := cpy.Py_fopen_obj(cpy.Py_BuildValue("s", "test.py"), "r+")
	//cpy.PyRun_SimpleFile(fp, "text.py")

	cpy.Py_Finalize()
}
