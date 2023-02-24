package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"strconv"
	"unsafe"
	"fmt"
)

type PyFloat struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyFloat) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyFloat) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyFloat) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyFloat) Str() string {
	return p.ToString()
}

func (p *PyFloat) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyFloat) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyFloat) Instance() uintptr {
	return p._instance()
}
func (p *PyFloat) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyFloat) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyFloat) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyFloat) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyFloat) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyFloat) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyFloat) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyFloat) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyFloat) AsInt() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyFloat) AsInt64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyFloat) AsUint() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyFloat) AsUint64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyFloat) AsFloat32() float32 {
	//暂时找不到syscall float实现
	//return cpy.PyLong_AsDouble(p._instance())
	v := cpy.PyObject_Str(p.Instance())
	f, _ := strconv.ParseFloat(v, 32)
	return float32(f)
}
func (p *PyFloat) AsDouble() float64 {
	//暂时找不到syscall float实现
	//return cpy.PyLong_AsDouble(p._instance())
	v := cpy.PyObject_Str(p.Instance())
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

// auto free
func NewPyFloatWithPtr(ptr uintptr) *PyFloat {
	o := new(PyFloat)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyLong).Free)
	return o
}

// Return value: New reference
func PyFloat_FromFloat64(v float64) *PyLong {
	u:=NewPyUnicode(fmt.Sprintf("%f",v))
	defer u.Free()
	return NewPyLongWithPtr(cpy.PyFloat_FromString(u.Instance()))
}

// Return value: New reference
func PyFloat_FromFloat32(v float32) *PyLong {
	u:=NewPyUnicode(fmt.Sprintf("%f",v))
	defer u.Free()
	return NewPyLongWithPtr(cpy.PyFloat_FromString(u.Instance()))
}

// Return value: New reference
func PyFloat_FromString(s string) *PyLong {
	u:=NewPyUnicode(s)
	defer u.Free()
	return NewPyLongWithPtr(cpy.PyFloat_FromString(u.Instance()))
}

func AsPyFloat(obj interface{}) *PyFloat {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyFloat{instance: instance}
}
