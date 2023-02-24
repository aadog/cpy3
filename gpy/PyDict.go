package gpy

import (
	"github.com/aadog/cpy3/cpy"
	"unsafe"
)

type PyDict struct {
	IPyObject
	instance unsafe.Pointer
}

func (p *PyClass) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyClass) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

// Return value: New reference
func (p *PyDict) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyDict) Str() string {
	return p.ToString()
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
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
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

// Return value: Borrowed reference.
func (p *PyDict) GetItemString(key string) *PyObject {
	return AsPyObject(cpy.PyDict_GetItemString(p._instance(), key))
}
func (p *PyDict) GetItemStringToStringOr(key string,def string)string{
	item:=p.GetItemString(key)
	if item.Str()=="None"{
		return def
	}
	return AsPyUnicode(item).AsUTF8()
}
func (p *PyDict) GetItemStringToFloat32Or(key string,def float32)float32{
	item:=p.GetItemString(key)
	if item.Str()=="None"{
		return def
	}
	return AsPyFloat(item).AsFloat32()
}
func (p *PyDict) GetItemStringToDoubleOr(key string,def float64)float64{
	item:=p.GetItemString(key)
	if item.Str()=="None"{
		return def
	}
	return AsPyLong(item).AsDouble()
}

func (p *PyDict) GetItemStringToIntOr(key string,def int)int{
	item:=p.GetItemString(key)
	if item.Str()=="None"{
		return def
	}
	return AsPyLong(item).AsInt()
}
func (p *PyDict) GetItemStringToInt64Or(key string,def int64)int64{
	item:=p.GetItemString(key)
	if item.Str()=="None"{
		return def
	}
	return AsPyLong(item).AsInt64()
}

// Return value: New reference.
func (p *PyDict) Keys() *PyObject {
	return NewPyObjectWithPtr(cpy.PyDict_Keys(p._instance()))
}

// Return value: Borrowed reference.
func (p *PyDict) KeysToStrings() []string {
	keys:=AsPyList(NewPyObjectWithPtr(cpy.PyDict_Keys(p._instance())))
	defer keys.Free()
	results:=make([]string,0)
	for i := int64(0); i < keys.Size(); i++ {
		results = append(results, keys.GetItem(i).Str())
	}
	return results
}

// Return value: Borrowed reference.
func (p *PyDict) GetItem(key *PyObject) *PyObject {
	return AsPyObject(cpy.PyDict_GetItem(p._instance(), key._instance()))
}

func (p *PyDict) SetItemString(key string, val IPyObject) int {
	return cpy.PyDict_SetItemString(p._instance(), key, val.Instance())
}
func (p *PyDict) SetItem(key IPyObject, val IPyObject) int {
	return cpy.PyDict_SetItem(p._instance(), key.Instance(), val.Instance())
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
	//setFinalizer(o, (*PyDict).Free)
	return o
}

func AsPyDict(obj interface{}) *PyDict {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	return &PyDict{instance: instance}
}
// Return value: New reference.
func NewPyDict() *PyDict {
	return NewPyDictWithPtr(cpy.PyDict_New())
}
