package main

import (
	_ "github.com/aadog/cpy3/alib"
	"github.com/aadog/cpy3/gpy"
	"os"
	"runtime"
)

/*
extern void* GoPyModule();
 */
import "C"
import "unsafe"

//export GoPyModule
func GoPyModule()uintptr{
	GoModule := gpy.CreateGoModule("_ui", "bb")
	GoModule.AddFunction("T", func() []byte {
		return make([]byte, 10240)
	})
	return GoModule.Instance()
}

func main1() {

	gpy.SetProgramName(os.Args[0])
	gpy.SetPythonHome("./")
	gpy.PyImport_AppendInittab("_ui",uintptr(unsafe.Pointer(&C.GoPyModule)))
	gpy.Initialize()

	gpy.RunSimpleString(
		`
import _ui
import time

for i in range(0,1000000):
	_ui.Call('T')
	time.sleep(0.01)
	`)
	runtime.GC()
	gpy.Finalize()
}

func main() {
	main1()
	//
	//main1()
}
