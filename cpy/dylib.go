//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cpy

import (
	dllimports2 "github.com/aadog/cpy3/cpy/dllimports"
	"runtime"
	"unsafe"

	"github.com/aadog/cpy3/pkgs/libname"
)

var (

	// 全局导入库
	uiLib            = loadUILib()
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func loadUILib() dllimports2.DLL {
	libName := getDLLName()
	// 如果支持运行时释放，则使用此种方法
	if support, newDLLPath := checkAndReleaseDLL(); support {
		libName = newDLLPath
	} else {
		if libname.LibName != "" {
			libName = libname.LibName
		}
	}
	lib, err := dllimports2.NewDLL(libName)
	if err != nil {
		panic(err)
	}
	return lib
}

func closeLib() {
	if uiLib != 0 {
		uiLib.Release()
		uiLib = 0
	}
}

func getDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

// 调用自动生成的API列表中的函数
func syscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports2.GetImportFunc(uiLib, trap).Call(args...)
	return r1
}

func syscallGetTextBuf(trap int, obj uintptr, buffer *string, bufSize uintptr) uintptr {
	if buffer == nil || bufSize == 0 {
		return 0
	}
	strPtr := make([]uint8, bufSize+1)
	sLen := syscallN(trap, obj, uintptr(unsafe.Pointer(&strPtr[0])), bufSize)
	if sLen > 0 {
		*buffer = string(strPtr[:sLen])
	}
	return sLen
}

// 调用手动导入的API列表中的函数
func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports2.GetImportDefFunc(uiLib, trap).Call(args...)
	return r1
}
