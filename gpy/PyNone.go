package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyNone struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyNone) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyNone) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyNone) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyNone) Str() string {
	return p.ToString()
}

func (p *PyNone) DecRef() {
	cpy.Py_DecRef(p._instance())
}
func (p *PyNone) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyNone) Instance() uintptr {
	return p._instance()
}
func (p *PyNone) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyNone) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyNone) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyNone) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyNone) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyNone) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyNone) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyNone) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

// auto free
func NewPyNoneWithPtr(ptr uintptr) *PyNone {
	o := new(PyNone)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyNone).Free)
	return o
}

func PyReturnNone() *PyNone {
	none := NewPyNoneWithPtr(cpy.Py_None())
	cpy.Py_IncRef(none._instance())
	return none
}
