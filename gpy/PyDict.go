package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyDict struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyDict) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyDict) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyDict) Instance() uintptr {
	return p._instance()
}
func (p *PyDict) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyDict) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyDict) ClassName() string {
	return PyObjectType(p).ClassName()
}

func (p *PyDict) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyDict) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyDict) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyDict) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyDict) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyDict) DelItemString(key string) int {
	return cpy.PyDict_DelItemString(p._instance(), key)
}

func (p *PyDict) DelItem(key *PyObject) int {
	return cpy.PyDict_DelItem(p._instance(), key._instance())
}

func (p *PyDict) GetItemString(key string) *PyObject {
	return AsPyObject(cpy.PyDict_GetItemString(p._instance(), key))
}

func (p *PyDict) Keys() *PyObject {
	return NewPyObjectWithPtr(cpy.PyDict_Keys(p._instance()))
}
func (p *PyDict) GetItem(key *PyObject) *PyObject {
	return AsPyObject(cpy.PyDict_GetItem(p._instance(), key._instance()))
}

func (p *PyDict) SetItemString(key string, val *PyObject) int {
	return cpy.PyDict_SetItemString(p._instance(), key, val._instance())
}
func (p *PyDict) SetItem(key *PyObject, val *PyObject) int {
	return cpy.PyDict_SetItem(p._instance(), key._instance(), val._instance())
}

func (p *PyDict) Size() int64 {
	return cpy.PyDict_Size(p._instance())
}
func (p *PyDict) Clear() {
	cpy.PyDict_Clear(p._instance())
}

// auto free
func NewPyDictWithPtr(ptr uintptr) *PyDict {
	o := new(PyDict)
	o.instance = unsafe.Pointer(ptr)
	setFinalizer(o, (*PyDict).Free)
	return o
}

func AsPyDict(obj interface{}) *PyDict {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyDict{instance: instance}
}

func NewPyDict() *PyDict {
	return NewPyDictWithPtr(cpy.PyDict_New())
}
