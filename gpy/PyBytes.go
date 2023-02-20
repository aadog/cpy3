package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyBytes struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyBytes) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyBytes) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyBytes) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyBytes) Str() string {
	return p.ToString()
}

func (p *PyBytes) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyBytes) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyBytes) Instance() uintptr {
	return p._instance()
}
func (p *PyBytes) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyBytes) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyBytes) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyBytes) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyBytes) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyBytes) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyBytes) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyBytes) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func AsPyBytes(obj interface{}) *PyBytes {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyBytes{instance: instance}
}

// auto free
func NewPyBytesWithPtr(ptr uintptr) *PyBytes {
	o := new(PyBytes)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyBytes).Free)
	return o
}

func PyBytes_FromString(s string) *PyBytes {
	return NewPyBytesWithPtr(cpy.PyBytes_FromString(s))
}
