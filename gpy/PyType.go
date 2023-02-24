package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyType struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyType) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyType) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyType) Name() string {
	return p.ClassName()
}

func (p *PyType) Str() string {
	return p.ToString()
}

func (p *PyType) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyType) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyType) Instance() uintptr {
	return p._instance()
}

func (p *PyType) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyType) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

// Return value: Borrowed reference.
func (p *PyType) ClassName() string {
	obj := cpy.PyObject_GetAttrString(p._instance(), "__name__")
	defer cpy.Py_DecRef(obj)
	return cpy.PyUnicode_AsUTF8(obj)
}

func (p *PyType) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyType) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyType) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

// auto free
func NewPyTypeWithPtr(ptr uintptr) *PyType {
	o := new(PyType)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyType).Free)
	return o
}

// Return value: New reference
func PyObjectType(owner IPyObject) *PyType {
	return NewPyTypeWithPtr(cpy.PyObject_Type(CheckPtr(owner)))
}

// Return value: Borrowed reference.
func PyObjectTypeName(owner IPyObject)string{
	tp:=PyObjectType(owner)
	defer tp.Free()
	return tp.ClassName()
}

func AsType(obj interface{}) *PyType {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyType{instance: instance}
}
