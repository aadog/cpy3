package gpy

import (
	"reflect"
)

// CheckPtr
//
// 检测接口是否被实例化，如果已经实例化则返回实例指针。
//
// Checks if the interface is instantiated, and returns an instance pointer if it has been instantiated.
//
//go:noinline
func CheckPtr(value interface{}) uintptr {
	switch value.(type) {
	case IPyObject:
		if reflect.ValueOf(value).Pointer() == 0 {
			return 0
		}
		return value.(IPyObject).Instance()
	}
	return 0
}
