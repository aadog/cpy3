package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyModule struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyModule) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyModule) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyModule) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyModule) Str() string {
	return p.ToString()
}

func (p *PyModule) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyModule) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyModule) Instance() uintptr {
	return p._instance()
}
func (p *PyModule) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyModule) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyModule) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyModule) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyModule) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyModule) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyModule) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyModule) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyModule) GetName() string {
	return cpy.PyModule_GetName(p._instance())
}
func (p *PyModule) GetDict() *PyDict {
	return AsPyDict(cpy.PyModule_GetDict(p._instance()))
}

//	func (p *PyModule) AddClass(class *PyClass) {
//		className := class.GetAttrString("__className__")
//		p.AddObject(className.AsUTF8(), class.AsObj())
//	}
func (p *PyModule) AddIntConstant(name string, value int64) int {
	return cpy.PyModule_AddIntConstant(p._instance(), name, value)
}
func (p *PyModule) AddStringConstant(name string, value string) int {
	return cpy.PyModule_AddStringConstant(p._instance(), name, value)
}
func (p *PyModule) AddObject(name string, value *PyObject) int {
	return cpy.PyModule_AddObject(p._instance(), name, value._instance())
}
func (p *PyModule) AddObjectRef(name string, value *PyObject) int {
	return cpy.PyModule_AddObjectRef(p._instance(), name, value._instance())
}
func AsPyModule(obj interface{}) *PyModule {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	o := &PyModule{instance: instance}
	return o
}
