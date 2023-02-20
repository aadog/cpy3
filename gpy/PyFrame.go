package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyFrame struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyFrame) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyFrame) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyFrame) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyFrame) Str() string {
	return p.ToString()
}

func (p *PyFrame) DecRef() {
	cpy.Py_DecRef(p._instance())
}
func (p *PyFrame) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyFrame) Instance() uintptr {
	return p._instance()
}
func (p *PyFrame) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyFrame) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyFrame) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyFrame) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyFrame) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyFrame) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyFrame) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyFrame) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyFrame) GetBack() *PyFrame {
	return NewPyFrameWithPtr(cpy.PyFrame_GetBack(p._instance()))
}
func (p *PyFrame) GetCode() *PyFrame {
	return NewPyFrameWithPtr(cpy.PyFrame_GetCode(p._instance()))
}

// auto free
func NewPyFrameWithPtr(ptr uintptr) *PyFrame {
	o := new(PyFrame)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyFrame).Free)
	return o
}

func AsPyFrame(obj interface{}) *PyFrame {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyFrame{instance: instance}
}
