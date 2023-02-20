package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyObject struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyObject) Str() string {
	return p.ToString()
}

func (p *PyObject) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func AsPyNil() *PyObject {
	return &PyObject{instance: unsafe.Pointer(uintptr(0))}
}

func AsPyObject(obj interface{}) *PyObject {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyObject{instance: instance}
}

func (p *PyObject) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}
func (p *PyObject) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyObject) Instance() uintptr {
	return p._instance()
}

func (p *PyObject) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyObject) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyObject) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyObject) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyObject) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyObject) CallNoArgs() IPyObject {
	return AsPyObject(cpy.PyObject_CallNoArgs(p.Instance()))
}
func (p *PyObject) Call(args IPyObject, kwargs IPyObject) IPyObject {
	return AsPyObject(cpy.PyObject_Call(p.Instance(), args.Instance(), kwargs.Instance()))
}
func (p *PyObject) CallObject(args IPyObject) IPyObject {
	return AsPyObject(cpy.PyObject_CallObject(p.Instance(), args.Instance()))
}
func (p *PyObject) PyObject_Call(args IPyObject, kwargs IPyObject) IPyObject {
	return AsPyObject(cpy.PyObject_Call(p.Instance(), args.Instance(), kwargs.Instance()))
}

func (p *PyObject) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyObject) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

// Get class type information.
func (p *PyObject) Type() *PyType {
	return PyObjectType(p)
}

// auto free
func NewPyObjectWithPtr(ptr uintptr) *PyObject {
	p := new(PyObject)
	p.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyObject).Free)
	return p
}
