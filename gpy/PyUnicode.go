package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyUnicode struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyUnicode) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyUnicode) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (o *PyUnicode) Type() *PyType {
	return PyObjectType(o)
}

func (p *PyUnicode) Str() string {
	return p.ToString()
}

func (p *PyUnicode) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyUnicode) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyUnicode) Instance() uintptr {
	return p._instance()
}
func (p *PyUnicode) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyUnicode) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyUnicode) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyUnicode) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyUnicode) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyUnicode) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyUnicode) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyUnicode) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyUnicode) GetLength() int64 {
	return cpy.PyUnicode_GetLength(p._instance())
}

func (p *PyUnicode) AsUTF8() string {
	return cpy.PyUnicode_AsUTF8(p._instance())
}

func PyUnicode_DecodeFSDefault(s string) *PyUnicode {
	return AsPyUnicode(cpy.PyUnicode_DecodeFSDefault(cpy.PascalStr(s)))
}

// auto free
func NewPyUnicodeWithPtr(ptr uintptr) *PyUnicode {
	o := new(PyUnicode)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyUnicode).Free)
	return o
}

func NewPyUnicode(s string) *PyUnicode {
	return NewPyUnicodeWithPtr(cpy.PyUnicode_FromString(s))
}

func AsPyUnicode(obj interface{}) *PyUnicode {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyUnicode{instance: instance}
}
