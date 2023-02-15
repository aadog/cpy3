package gpy

import (
	"fmt"
	"github.com/aadog/cpy3/cpy"
	"reflect"
	"sync"
	"syscall"
	"unsafe"
)

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

type PyModule struct {
	IPyObject
	GoObj    *PyModuleGoObj
	instance unsafe.Pointer
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
	return PyObjectType(p).ClassName()
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

func (p *PyModule) AddFunction(name string, fn interface{}) {
	p.GoObj.CallMap.Store(name, fn)
}
func (p *PyModule) AddClass(class *PyClass) {
	className := class.GetAttrString("__className__")
	p.AddObject(className.AsUTF8(), class.AsObj())
}
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

// auto free
func NewPyModuleWithPtr(ptr uintptr) *PyModule {
	o := new(PyModule)
	o.instance = unsafe.Pointer(ptr)
	name := o.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	o.GoObj = smmodule.(*PyModule).GoObj
	setFinalizer(o, (*PyModule).Free)
	return o
}

func AsPyModule(obj interface{}) *PyModule {
	instance := getInstance(obj)
	if instance == nullptr {
		return nil
	}
	o := &PyModule{instance: instance}
	name := o.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	o.GoObj = smmodule.(*PyModule).GoObj
	return o
}

func PyTypeToGoType(p *PyObject) any {

	return 0
}

func PyMethodForward(self *PyModule, args *PyTuple, method interface{}) *PyObject {

	methodType := reflect.TypeOf(method)
	methodValue := reflect.ValueOf(method)
	if methodType.Kind() != reflect.Func {
		return Py_RETURN_NONE().AsObj()
	}
	if int64(methodType.NumIn()) != args.Size() {
		PyErr_SetString(UserException(), fmt.Sprintf("The number of parameters does not match,%d parameter is required, and you have entered %d", methodType.NumIn(), args.Size()))
		return Py_RETURN_NONE().AsObj()
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
			defer obj.DecRef()
			l.Append(obj)
		}
		return l.AsObj()
	}
	return Py_RETURN_NONE().AsObj()
}

var PyModuleMethodForwardCallBack = syscall.NewCallback(func(self uintptr, args uintptr) uintptr {

	pyArgs := PyTupleFromInst(args)
	pyArgsLen := pyArgs.Size()
	if pyArgsLen < 1 {
		return Py_RETURN_NONE().instance
	}
	ForwardCode := pyArgs.GetItem(0).Str()

	pyModule := PyModuleFromInst(self)
	moduleName := pyModule.GetName()
	ifn, ok := pyModule.GoObj.CallMap.Load(ForwardCode)
	if ok == false {
		PyErr_SetString(UserException(), fmt.Sprintf("%s not find method %s ", moduleName, ForwardCode))
		return Py_RETURN_NONE().Instance()
	}
	//处理参数
	newArgs := PyTupleFromObj(pyArgs.GetSlice(1, pyArgsLen))
	defer newArgs.DecRef()

	return PyMethodForward(pyModule, newArgs, ifn).Instance()
})

func CreateModule(name string, doc string) *PyModule {
	pyModuleMethodCallDefs := make([]cpy3.PyMethodDef, 0)
	methodCallDef := cpy3.PyMethodDef{
		Ml_name:  msvcrt.MallocCString("Call"),
		Ml_meth:  PyModuleMethodForwardCallBack,
		Ml_flags: 1,
		Ml_doc:   msvcrt.MallocCString("module call forward"),
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, methodCallDef)
	moduleNullMethodDef := cpy3.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, moduleNullMethodDef)

	module := &PyModule{}
	module.GoObj = new(PyModuleGoObj)
	module.GoObj.methodsDef = pyModuleMethodCallDefs
	module.GoObj.moduleDef = &cpy3.PyModuleDef{
		M_base: cpy3.PyModuleDef_Base{
			Ob_base: cpy3.PyObject_HEAD_INIT(0),
		},
		M_name:     msvcrt.MallocCString(name),
		M_doc:      msvcrt.MallocCString(doc),
		M_size:     -1,
		M_methods:  uintptr(unsafe.Pointer(&pyModuleMethodCallDefs[0])),
		M_slots:    0,
		M_traverse: 0,
		M_clear:    0,
		M_free:     0,
	}
	ptr := cpy3.PyModule_Create2(uintptr(unsafe.Pointer(module.GoObj.moduleDef)), PYTHON_API_VERSION)
	module.instance = ptr
	module.ptr = unsafe.Pointer(module.instance)
	SystemModuleMap.Store(name, module)
	return module
}

func NewModuleInitFuncCallBack(moduleName string, fn func() IPyObject) uintptr {
	var c = syscall.NewCallback(func() uintptr {
		return fn().Instance()
	})
	pyModuleInitMap.Store(moduleName, c)
	return c
}
