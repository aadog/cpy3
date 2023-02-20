package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyTuple struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyTuple) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyTuple) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (o *PyTuple) Type() *PyType {
	return PyObjectType(o)
}

func (p *PyTuple) Str() string {
	return p.ToString()
}

func (p *PyTuple) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyTuple) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyTuple) Instance() uintptr {
	return p._instance()
}
func (p *PyTuple) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyTuple) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyTuple) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyTuple) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyTuple) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyTuple) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyTuple) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyTuple) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyTuple) Size() int64 {
	return cpy.PyTuple_Size(p._instance())
}
func (p *PyTuple) Check() int64 {
	return cpy.PyTuple_Check(p._instance())
}
func (p *PyTuple) GetItem(pos int64) *PyObject {
	return AsPyObject(cpy.PyTuple_GetItem(p._instance(), pos))
}
func (p *PyTuple) SetItem(pos int64, o IPyObject) int {
	return cpy.PyTuple_SetItem(p._instance(), pos, o.Instance())
}
func (p *PyTuple) GetSlice(low int64, high int64) *PyList {
	return NewPyListWithPtr(cpy.PyTuple_GetSlice(p.Instance(), low, high))
}

// auto free
func NewPyTupleWithPtr(ptr uintptr) *PyTuple {
	o := new(PyTuple)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyTuple).Free)
	return o
}

func AsPyTuple(obj interface{}) *PyTuple {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyTuple{instance: instance}
}

func NewPyTuple(l int64) *PyTuple {
	return NewPyTupleWithPtr(cpy.PyTuple_New(l))
}
