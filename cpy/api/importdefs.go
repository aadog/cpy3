package api

import (
	"github.com/aadog/cpy3/cpy/api/dllimports"
	"unsafe"
)

func Py_Initialize() {
	defSyscallN(dllimports.Py_Initialize)
}

func Py_InitializeEx(initsigs int) {
	defSyscallN(dllimports.Py_InitializeEx, uintptr(initsigs))
}
func Py_IsInitialized() int {
	return int(defSyscallN(dllimports.Py_IsInitialized))
}
func Py_Finalize() {
	defSyscallN(dllimports.Py_Finalize)
}

func Py_FinalizeEx() int {
	return int(defSyscallN(dllimports.Py_FinalizeEx))
}

// 使用PyMem_RawFree释放
func StringToLocale(s string) uintptr {
	p := Py_DecodeLocale(s, nil)
	return p
}

// to string
func PyLocaleToString(obj uintptr, len int) string {
	p := Py_EncodeLocale(obj, uintptr(0))
	defer PyMem_Free(p)
	return GoStr1(p, len)
}

// 使用PyMem_Free释放
func Py_EncodeLocale(obj uintptr, error_pos uintptr) uintptr {
	return defSyscallN(dllimports.Py_EncodeLocale, obj, error_pos)
}

// 使用PyMem_RawFree释放
func Py_DecodeLocale(arg string, size *int) uintptr {
	return defSyscallN(dllimports.Py_DecodeLocale, PascalStr(arg), uintptr(unsafe.Pointer(size)))
}

func PyMem_Free(obj uintptr) {
	defSyscallN(dllimports.PyMem_Free, obj)
}

func PyMem_RawFree(obj uintptr) {
	defSyscallN(dllimports.PyMem_RawFree, obj)
}

func Py_Main(args []string) int {
	argc := len(args)
	argv := make([]uintptr, 0)
	for _, arg := range args {
		argv = append(argv, PascalStr(arg))
	}
	r := defSyscallN(dllimports.Py_Main, uintptr(argc), uintptr(unsafe.Pointer(&argv[0])))
	return int(r)
}

func Py_BytesMain(args []string) int {
	argc := len(args)
	argv := make([]uintptr, 0)
	for _, arg := range args {
		argv = append(argv, PascalStr(arg))
	}
	r := defSyscallN(dllimports.Py_Main, uintptr(argc), uintptr(unsafe.Pointer(&argv[0])))
	return int(r)
}

func PyRun_AnyFile(fp uintptr, filename string) int {
	r := defSyscallN(dllimports.PyRun_AnyFile, uintptr(fp), PascalStr(filename))
	return int(r)
}
func PyRun_SimpleFile(fp uintptr, filename string) int {
	r := defSyscallN(dllimports.PyRun_SimpleFile, uintptr(fp), PascalStr(filename))
	return int(r)
}
func PyRun_SimpleString(command string) int {
	r := defSyscallN(dllimports.PyRun_SimpleString, PascalStr(command))
	return int(r)
}

func Py_SetProgramName(name string) {
	p := StringToLocale(name)
	defSyscallN(dllimports.Py_SetProgramName, p)
}
func Py_GetProgramName() string {
	r := defSyscallN(dllimports.Py_GetProgramName)
	return Utf16ToGoStr(r)
}

func Py_SetPath(path string) {
	p := StringToLocale(path)
	defer PyMem_RawFree(p)
	defSyscallN(dllimports.Py_SetPath, p)
}

func Py_GetPath() string {
	r := defSyscallN(dllimports.Py_GetPath)
	return Utf16ToGoStr(r)
}

func Py_GetPythonHome() string {
	r := defSyscallN(dllimports.Py_GetPythonHome)
	return Utf16ToGoStr(r)
}

func Py_SetPythonHome(home string) {
	p := StringToLocale(home)
	defSyscallN(dllimports.Py_SetPythonHome, p)
}

//func Py_fopen_obj(path uintptr, mode string) uintptr {
//	r, _, _ := _py_fopen_obj.Call(path, msvcrt.StringToCUTF8String(mode))
//	r:=defSyscallN(dllimports.py_)
//	return r
//}

func Py_IncRef(obj uintptr) {
	defSyscallN(dllimports.Py_IncRef, obj)
}
func Py_DecRef(obj uintptr) {
	defSyscallN(dllimports.Py_DecRef, obj)
}

func PyImport_AppendInittab(name string, initfunc uintptr) int {
	r := defSyscallN(dllimports.PyImport_AppendInittab, PascalStr(name), initfunc)
	return int(r)
}
