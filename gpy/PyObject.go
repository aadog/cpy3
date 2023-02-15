package gpy

import (
	"fmt"
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyObject struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyObject) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func AsPyObject(obj interface{}) *PyObject {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyObject{instance: instance}
}

func (o *PyObject) Free() {
	if o.instance != nullptr {
		fmt.Println("不知道为什么有一个错误")
		//cpy.Py_DecRef(o._instance())
		//cpy.PyObject_Del(o.Instance())
		o.instance = nullptr
	}
}
func (o *PyObject) _instance() uintptr {
	return uintptr(o.instance)
}

func (o *PyObject) Instance() uintptr {
	return o._instance()
}

func (o *PyObject) UnsafeAddr() unsafe.Pointer {
	return o.instance
}

func (o *PyObject) IsValid() bool {
	return o.instance != nullptr
}

func (o *PyObject) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(o._instance()))
}

func (o *PyObject) ToString() string {
	return cpy.PyObject_Str(o._instance())
}

func (o *PyObject) ClassName() string {
	return PyObjectType(o).ClassName()
}

//// Get class type information.
//func (o *TObject) ClassType() TClass {
//	return Object_ClassType(o._instance())
//}

// auto free
func NewPyObjectWithPtr(ptr uintptr) *PyObject {
	o := new(PyObject)
	o.instance = unsafe.Pointer(ptr)
	setFinalizer(o, (*PyObject).Free)
	return o
}
