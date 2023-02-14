package main

import (
	"fmt"
	_ "github.com/aadog/cpy3/alib"
	"github.com/aadog/cpy3/cpy/api"
)

func main() {

	api.Py_Initialize()
	u := api.PyUnicode_FromString("ff")
	fmt.Println(api.PyUnicode_AsUTF8(u))
	//api.Py_SetProgramName("xxxx")
	//api.Py_SetPath("aaa")
	//api.Py_SetPythonHome("bb")
	fmt.Println(api.Py_GetProgramName())
	fmt.Println(api.Py_GetPath())
	fmt.Println(api.Py_GetPythonHome())
	api.Py_Finalize()
}
