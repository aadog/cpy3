package gpy

import (
	"fmt"
	"github.com/aadog/cpy3/cpy"
	"reflect"
	"sync"
	"unsafe"
)

/*
extern void* doPyGoModuleMethodForwardCallBack(void* self ,void* args);
 */
import "C"

//export doPyGoModuleMethodForwardCallBack
func doPyGoModuleMethodForwardCallBack(self unsafe.Pointer,args unsafe.Pointer)unsafe.Pointer{
	pyArgs := AsPyTuple(args)
	pyArgsLen := pyArgs.Size()
	if pyArgsLen < 1 {
		return PyReturnNone().UnsafeAddr()
	}
	ForwardCode := pyArgs.GetItem(0).Str()

	pyModule := AsPyGoModule(self)
	moduleName := pyModule.GetName()
	ifn, ok := pyModule.GoObj.CallMap.Load(ForwardCode)
	if ok == false {
		PyErr_SetString(UserException(), fmt.Sprintf("%s not find method %s ", moduleName, ForwardCode))
		return PyReturnNone().UnsafeAddr()
	}
	//处理参数
	newArgs := AsPyTuple(pyArgs.GetSlice(1, pyArgsLen))
	defer newArgs.Free()


	r:=PyMethodForward(pyModule, newArgs, ifn)
	if PyErr_Occurred()!=nil{
		return nil
	}
	return r.UnsafeAddr()
}

const (
	PYTHON_API_VERSION = 1013
	PYTHON_API_STRING  = "1013"
)

var pyModuleInitMap = sync.Map{}

type PyModuleGoObj struct {
	moduleDef  *cpy.PyModuleDef
	methodsDef []cpy.PyMethodDef
	CallMap    sync.Map
}

type PyGoModule struct {
	IPyObject
	GoObj    *PyModuleGoObj
	instance unsafe.Pointer
}

func (p *PyGoModule) IncRef() {
	if p._instance() != 0 {
		cpy.Py_IncRef(p._instance())
	}
}
func (p *PyGoModule) RefCount() int {
	return int(cpy.PyObjectFromPtr(p._instance()).Ob_refcnt)
}

func (p *PyGoModule) Type() *PyType {
	return PyObjectType(p)
}

func (p *PyGoModule) Str() string {
	return p.ToString()
}

func (p *PyGoModule) DecRef() {
	cpy.Py_DecRef(p._instance())
}

func (p *PyGoModule) _instance() uintptr {
	return uintptr(p.instance)
}

func (p *PyGoModule) Instance() uintptr {
	return p._instance()
}
func (p *PyGoModule) IsValid() bool {
	return p.instance != nullptr
}

func (p *PyGoModule) UnsafeAddr() unsafe.Pointer {
	return p.instance
}

func (p *PyGoModule) ClassName() string {
	tp := PyObjectType(p)
	defer tp.Free()
	return tp.ClassName()
}

func (p *PyGoModule) Free() {
	if p.instance != nullptr {
		cpy.Py_DecRef(p._instance())
		p.instance = nullptr
	}
}

func (p *PyGoModule) GetHashCode() int64 {
	return int64(cpy.PyObject_Hash(p._instance()))
}

func (p *PyGoModule) Equals(object IPyObject) bool {
	//TODO implement me
	panic("implement me")
}

func (p *PyGoModule) InstanceSize() int32 {
	//TODO implement me
	panic("implement me")
}

func (p *PyGoModule) ToString() string {
	return cpy.PyObject_Str(p._instance())
}

func (p *PyGoModule) GetName() string {
	return cpy.PyModule_GetName(p._instance())
}
func (p *PyGoModule) GetDict() *PyDict {
	return AsPyDict(cpy.PyModule_GetDict(p._instance()))
}

func (p *PyGoModule) AddFunction(name string, fn interface{}) {
	p.GoObj.CallMap.Store(name, fn)
}

//	func (p *PyModule) AddClass(class *PyClass) {
//		className := class.GetAttrString("__className__")
//		p.AddObject(className.AsUTF8(), class.AsObj())
//	}
func (p *PyGoModule) AddIntConstant(name string, value int64) int {
	return cpy.PyModule_AddIntConstant(p._instance(), name, value)
}
func (p *PyGoModule) AddStringConstant(name string, value string) int {
	return cpy.PyModule_AddStringConstant(p._instance(), name, value)
}
func (p *PyGoModule) AddObject(name string, value *PyObject) int {
	return cpy.PyModule_AddObject(p._instance(), name, value._instance())
}
func (p *PyGoModule) AddObjectRef(name string, value *PyObject) int {
	return cpy.PyModule_AddObjectRef(p._instance(), name, value._instance())
}

// auto free
func NewPyGoModuleWithPtr(ptr uintptr) *PyGoModule {
	o := new(PyGoModule)
	o.instance = unsafe.Pointer(ptr)
	name := o.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	o.GoObj = smmodule.(*PyGoModule).GoObj
	setFinalizer(o, (*PyModule).Free)
	return o
}

func AsPyGoModule(obj interface{}) *PyGoModule {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	o := &PyGoModule{instance: instance}
	name := o.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	o.GoObj = smmodule.(*PyGoModule).GoObj
	return o
}

func PyMethodForward(self *PyGoModule, args *PyTuple, method interface{}) IPyObject {
	methodType := reflect.TypeOf(method)
	methodValue := reflect.ValueOf(method)
	if methodType.Kind() != reflect.Func {
		return AsPyObject(PyReturnNone())
	}
	if int64(methodType.NumIn()) != args.Size() {
		PyErr_SetString(UserException(), fmt.Sprintf("The number of parameters does not match,%d parameter is required, and you have entered %d", methodType.NumIn(), args.Size()))
		return nil
	}

	fnArgs := make([]reflect.Value, 0)
	for i := 0; i < methodType.NumIn(); i++ {
		fnArgType := methodType.In(i)
		inArg := args.GetItem(int64(i))
		fnArgs = append(fnArgs, reflect.ValueOf(PyObjectToGo(inArg, fnArgType)))
	}
	rets := methodValue.Call(fnArgs)
	if len(rets) == 1 {
		firstRet := rets[0]
		r := GoToPyObject(firstRet.Interface())
		return r
	} else if len(rets) > 1 {
		l := NewPyList(0)
		for _, r := range rets {
			obj := GoToPyObject(r.Interface())
			defer obj.Free()
			l.Append(obj)
		}
		return l
	}
	return AsPyObject(PyReturnNone())
}



func CreateGoModule(name string, doc string) *PyGoModule {
	pyModuleMethodCallDefs := make([]cpy.PyMethodDef, 0)
	methodCallDef := cpy.PyMethodDef{
		Ml_name:  cpy.PascalStr("Call"),
		Ml_meth:  uintptr(C.doPyGoModuleMethodForwardCallBack),
		Ml_flags: 1,
		Ml_doc:   cpy.PascalStr("module call forward"),
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, methodCallDef)
	moduleNullMethodDef := cpy.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, moduleNullMethodDef)

	module := &PyGoModule{}
	module.GoObj = new(PyModuleGoObj)
	module.GoObj.methodsDef = pyModuleMethodCallDefs
	module.GoObj.moduleDef = &cpy.PyModuleDef{
		M_base: cpy.PyModuleDef_Base{
			Ob_base: cpy.PyObject_HEAD_INIT(0),
		},
		M_name:     cpy.PascalStr(name),
		M_doc:      cpy.PascalStr(doc),
		M_size:     -1,
		M_methods:  uintptr(unsafe.Pointer(&pyModuleMethodCallDefs[0])),
		M_slots:    0,
		M_traverse: 0,
		M_clear:    0,
		M_free:     0,
	}
	ptr := cpy.PyModule_Create2(uintptr(unsafe.Pointer(module.GoObj.moduleDef)), PYTHON_API_VERSION)
	module.instance = unsafe.Pointer(ptr)
	SystemModuleMap.Store(name, module)
	return module
}

