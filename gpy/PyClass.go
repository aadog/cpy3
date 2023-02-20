package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"sync"
	"unsafe"
)

var SystemClassMap = sync.Map{}

type PyClassGoObj struct {
	CallMap sync.Map
}

type PyClass struct {
	IPyObject
	GoObj    *PyClassGoObj
	instance unsafe.Pointer
}

func (p *PyDict) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyDict) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyClass) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyClass) Str() string {
	return p.ToString()
}

func (p *PyClass) DecRef() {
	cpy.Py_DecRef(p._instance())
}
func (p *PyClass) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyClass) Instance() uintptr {
	return p._instance()
}
func (p *PyClass) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyClass) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyClass) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyClass) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyClass) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyClass) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyClass) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyClass) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyClass) GetAttrString(attr_name string) *PyObject {
	o := cpy.PyObject_GetAttrString(p._instance(), attr_name)
	return NewPyObjectWithPtr(o)
}

func (p *PyClass) GetName() string {
	return AsPyUnicode(p.GetAttrString("__className__")).AsUTF8()
}

func (p *PyClass) AddFunction(name string, fn interface{}) {
	p.GoObj.CallMap.Store(name, fn)
}

// auto free
func NewPyClassWithPtr(ptr uintptr) *PyClass {
	o := new(PyClass)
	o.instance = unsafe.Pointer(ptr)
	//setFinalizer(o, (*PyClass).Free)
	return o
}

func AsPyClass(obj interface{}) *PyClass {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	o := &PyClass{instance: instance}
	name := o.GetName()
	smmodule, _ := SystemClassMap.Load(name)
	o.GoObj = smmodule.(*PyClass).GoObj
	return o
}
