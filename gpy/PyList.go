package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyList struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyList) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyList) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyList) Instance() uintptr {
	return p._instance()
}
func (p *PyList) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyList) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyList) ClassName() string {
	return PyObjectType(p).ClassName()
}

func (p *PyList) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyList) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyList) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyList) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyList) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyList) SetItem(index int64, o *PyObject) int {
	return cpy.PyList_SetItem(p._instance(), index, o._instance())
}
func (p *PyList) GetItem(index int64) *PyObject {
	return AsPyObject(cpy.PyList_GetItem(p._instance(), index))
}
func (p *PyList) GetSlice(low int64, high int64) *PyObject {
	return AsPyObject(cpy.PyList_GetSlice(p._instance(), low, high))
}
func (p *PyList) Size() int64 {
	return cpy.PyList_Size(p._instance())
}
func (p *PyList) Insert(index int64, item *PyObject) int {
	return cpy.PyList_Insert(p._instance(), index, item._instance())
}
func (p *PyList) Append(item *PyObject) int {
	return cpy.PyList_Append(p._instance(), item._instance())
}

// auto free
func NewPyListWithPtr(ptr uintptr) *PyList {
	o := new(PyList)
	o.instance = unsafe.Pointer(ptr)
	setFinalizer(o, (*PyList).Free)
	return o
}

func AsPyList(obj interface{}) *PyList {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyList{instance: instance}
}

func NewPyList(len int64) *PyList {
	return NewPyListWithPtr(cpy.PyList_New(len))
}
