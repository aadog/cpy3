package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyLong struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyLong) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyLong) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyLong) Instance() uintptr {
	return p._instance()
}
func (p *PyLong) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyLong) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyLong) ClassName() string {
	return PyObjectType(p).ClassName()
}

func (p *PyLong) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyLong) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyLong) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyLong) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyLong) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyLong) AsInt() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyLong) AsInt64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyLong) AsUint() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyLong) AsUint64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyLong) AsFloat32() float32 {
	return float32(cpy.PyLong_AsDouble(p._instance()))
}
func (p *PyLong) AsDouble() float64 {
	return cpy.PyLong_AsDouble(p._instance())
}

// auto free
func NewPyLongWithPtr(ptr uintptr) *PyLong {
	o := new(PyLong)
	o.instance = unsafe.Pointer(ptr)
	setFinalizer(o, (*PyLong).Free)
	return o
}

func PyLongFromLong(n int) *PyLong {
	return NewPyLongWithPtr(cpy.PyLong_FromLong(n))
}

func PyLongFromLongLong(n int64) *PyLong {
	return NewPyLongWithPtr(cpy.PyLong_FromLongLong(n))
}

func PyLong_FromDouble(n float64) *PyLong {
	return NewPyLongWithPtr(cpy.PyLong_FromDouble(n))
}

func AsPyLong(obj interface{}) *PyLong {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyLong{instance: instance}
}
