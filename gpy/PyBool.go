package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyBool struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyBool) DecRef() {
	cpy.Py_DecRef(p._instance())
}
func (p *PyBool) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyBool) Instance() uintptr {
	return p._instance()
}
func (p *PyBool) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyBool) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyBool) ClassName() string {
	return PyObjectType(p).ClassName()
}

func (p *PyBool) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyBool) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyBool) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyBool) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyBool) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

// auto free
func NewPyBoolWithPtr(ptr uintptr) *PyBool {
	o := new(PyBool)
	o.instance = unsafe.Pointer(ptr)
	setFinalizer(o, (*PyBool).Free)
	return o
}

func AsPyBool(obj interface{}) *PyBool {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyBool{instance: instance}
}

func NewPyBool(b bool) *PyBool {
	return NewPyBoolWithPtr(cpy.PyBool_FromLong(b))
}
