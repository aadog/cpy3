package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"path/filepath"
)

func RunAnyFile(pathstr string) int {
	pathobj := NewPyUnicode(pathstr)
	defer pathobj.DecRef()
	fp := cpy.Py_fopen_obj(pathobj._instance(), "r")
	return cpy.PyRun_AnyFile(fp, filepath.Base(pathstr))
}
func RunSimpleFile(pathstr string) int {
	pathobj := NewPyUnicode(pathstr)
	defer pathobj.DecRef()
	fp := cpy.Py_fopen_obj(pathobj._instance(), "r")
	return cpy.PyRun_SimpleFile(fp, filepath.Base(pathstr))
}
func RunSimpleString(command string) int {
	return cpy.PyRun_SimpleString(command)
}
