package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"strconv"
	"unsafe"
	"fmt"
)

type PyLong struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyLong) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyLong) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyLong) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyLong) Str() string {
	return p.ToString()
}

func (p *PyLong) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyLong) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyLong) Instance() uintptr {
	return p._instance()
}
func (p *PyLong) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyLong) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyLong) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyLong) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyLong) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyLong) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyLong) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyLong) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyLong) AsInt() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyLong) AsInt64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyLong) AsUint() int {
	return cpy.PyLong_AsLong(p._instance())
}
func (p *PyLong) AsUint64() int64 {
	return cpy.PyLong_AsLongLong(p._instance())
}
func (p *PyLong) AsFloat32() float32 {
	//暂时找不到syscall float实现
	//return cpy.PyLong_AsDouble(p._instance())
	v := cpy.PyObject_Str(p.Instance())
	f, _ := strconv.ParseFloat(v, 64)
	return float32(f)
}
func (p *PyLong) AsDouble() float64 {
	//暂时找不到syscall float实现
	//return cpy.PyLong_AsDouble(p._instance())
	v := cpy.PyObject_Str(p.Instance())
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

// auto free
func NewPyLongWithPtr(ptr uintptr) *PyLong {
	o := new(PyLong)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyLong).Free)
	return o
}

func PyLongFromLong(n int) *PyLong {
	return NewPyLongWithPtr(cpy.PyLong_FromLong(n))
}

func PyLongFromLongLong(n int64) *PyLong {
	return NewPyLongWithPtr(cpy.PyLong_FromLongLong(n))
}

// Return value: New reference.
func PyLong_FromDouble(n float64) *PyLong {
	u:=NewPyUnicode(fmt.Sprintf("%f",n))
	defer u.Free()
	return NewPyLongWithPtr(cpy.PyLong_FromUnicodeObject(u.Instance(),10))
}
// Return value: New reference.
func PyLong_FromFloat32(n float32) *PyLong {
	//浮点传不过去用字符代替
	//u:=NewPyUnicode(fmt.Sprintf("%f",n))
	//defer u.Free()
	//fmt.Println(fmt.Sprintf("%f",n))
	//return NewPyLongWithPtr(cpy.PyLong_FromUnicodeObject(u.Instance(),10))
	return NewPyLongWithPtr(cpy.PyLong_FromVoidPtr(uintptr(unsafe.Pointer(&n))))
}

func AsPyLong(obj interface{}) *PyLong {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyLong{instance: instance}
}
