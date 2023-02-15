package gpy

import (
	"fmt"
	"github.com/aadog/cpy3/cpy"
	"sync"
	"syscall"
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
	return PyObjectType(p).ClassName()
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
	setFinalizer(o, (*PyClass).Free)
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

func PyClassInstanceMethodForward(self *PyObject, args *PyTuple, method interface{}) *PyObject {
	fmt.Println("class call")
	//	methodType := reflect.TypeOf(method)
	//	methodValue := reflect.ValueOf(method)
	//	if methodType.Kind() != reflect.Func {
	//		return Py_RETURN_NONE().AsObj()
	//	}
	//	if int64(methodType.NumIn()) != args.Size() {
	//		PyErr_SetString(UserException(), fmt.Sprintf("The number of parameters does not match,%d parameter is required, and you have entered %d", methodType.NumIn(), args.Size()))
	//		return Py_RETURN_NONE().AsObj()
	//	}
	//
	//	fnArgs := make([]reflect.Value, 0)
	//	for i := 0; i < methodType.NumIn(); i++ {
	//		fnArgType := methodType.In(i)
	//		inArg := args.GetItem(int64(i))
	//		fnArgs = append(fnArgs, reflect.ValueOf(PyObjectToGo(inArg, fnArgType)))
	//	}
	//	rets := methodValue.Call(fnArgs)
	//	if len(rets) == 1 {
	//		firstRet := rets[0]
	//		r := GoToPyObject(firstRet.Interface())
	//		return r
	//	} else if len(rets) > 1 {
	//		l := NewPyList(0)
	//		for _, r := range rets {
	//			obj := GoToPyObject(r.Interface())
	//			defer obj.DecRef()
	//			l.Append(obj)
	//		}
	//		return l.AsObj()
	//	}
	return Py_RETURN_NONE().AsObj()
}

func CreateClass(name string, doc string) *PyClass {
	class := &PyClass{}
	class.GoObj = new(PyClassGoObj)

	dict := map[string]any{}
	pClassName := GoToPyObject(name)
	defer pClassName.DecRef()
	pClassBases := NewPyTuple(0)
	defer pClassBases.DecRef()
	pClassDic := GoToPyObject(dict)
	defer pClassDic.DecRef()
	class.instance = cpy3.PyObject_CallFunctionObjArgs(cpy3.PyType_Type(), pClassName.Instance(), pClassBases.Instance(), pClassDic.Instance(), 0)
	class.ptr = unsafe.Pointer(class.instance)
	class.SetAttrString("__className__", pClassName)

	fn := PyObjectFromInst(cpy3.PyCFunction_New(uintptr(unsafe.Pointer(PyClassInstanceMethodCallDef)), PyNil.instance))
	defer fn.DecRef()
	method := PyObjectFromInst(cpy3.PyInstanceMethod(fn.instance))
	defer method.DecRef()
	class.SetAttrString("Call", method)

	SystemClassMap.Store(name, class)
	return class
}

var PyClassInstanceMethodCallDef *cpy3.PyMethodDef
var PyClassInstanceMethodForwardCallBack = syscall.NewCallback(func(self uintptr, args uintptr) uintptr {
	pyArgs := PyTupleFromInst(args)
	arg1 := pyArgs.GetItem(0)
	tp := arg1.Type()
	defer tp.DecRef()
	if tp.Name() == "str" {
		fmt.Println("static")
	} else {
		pyArgsLen := pyArgs.Size()
		if pyArgsLen < 2 {
			return Py_RETURN_NONE().instance
		}
		ForwardCode := pyArgs.GetItem(1).Str()

		//处理参数
		newArgs := PyTupleFromObj(pyArgs.GetSlice(2, pyArgsLen))
		defer newArgs.DecRef()
		pySelf := pyArgs.GetItem(0)

		PyClass := PyClassFromObj(pySelf)
		className := PyClass.GetName()
		ifn, ok := PyClass.GoObj.CallMap.Load(ForwardCode)
		if ok == false {
			PyErr_SetString(UserException(), fmt.Sprintf("%s not find method %s ", className, ForwardCode))
			return Py_RETURN_NONE().Instance()
		}

		return PyClassInstanceMethodForward(pySelf, newArgs, ifn).instance
	}
	return Py_RETURN_NONE().instance
})

//func init() {
//	PyClassInstanceMethodCallDef = &cpy3.PyMethodDef{
//		Ml_name:  cpy3.GoStrToCStr("Call"),
//		Ml_meth:  PyClassInstanceMethodForwardCallBack,
//		Ml_flags: 3,
//		Ml_doc:   cpy3.GoStrToCStr("class call forward"),
//	}
//}

//func CreateClass(name string, dict map[string]any) *PyClass {
//	if dict == nil {
//		dict = map[string]any{}
//	}
//	pClassName := GoToPyObject(name)
//	defer pClassName.DecRef()
//	pClassBases := NewPyTuple(0)
//	defer pClassBases.DecRef()
//	pClassDic := GoToPyObject(dict)
//	defer pClassDic.DecRef()
//	pClass := PyObjectFromInst(cpy3.PyObject_CallFunctionObjArgs(cpy3.PyType_Type(), pClassName.Instance(), pClassBases.Instance(), pClassDic.Instance(), 0))
//	pClass.SetAttrString("__className__", NewPyUnicode(name).AsObj())
//
//	fnName := "Call"
//	methodCallDef := &cpy3.PyMethodDef{
//		Ml_name: cpy3.GoStrToCStr(fnName),
//		Ml_meth: syscall.NewCallback(func(self uintptr, args uintptr) uintptr {
//			pyArgs := PyTupleFromInst(args)
//			pyArgsLen := pyArgs.Size()
//			if pyArgsLen < 2 {
//				return Py_RETURN_NONE().instance
//			}
//			pySelf := pyArgs.GetItem(0)
//			ForwardCode := pyArgs.GetItem(1).Str()
//
//			ifn, ok := PyMethodMap.Load(ForwardCode)
//			if ok == false {
//				return Py_RETURN_NONE().Instance()
//			}
//			//处理参数
//			newArgs := PyTupleFromObj(pyArgs.GetSlice(1, pyArgsLen))
//			defer newArgs.DecRef()
//			return PyInstanceMethodForward(pySelf, newArgs, ifn).Instance()
//		}),
//		Ml_flags: 3,
//		Ml_doc:   cpy3.GoStrToCStr("module call forward"),
//	}
//	PyMethodMap.Store("x", methodCallDef)
//
//	fn := PyObjectFromInst(cpy3.PyCFunction_New(uintptr(unsafe.Pointer(methodCallDef)), PyNil.instance))
//	defer fn.DecRef()
//	method := PyObjectFromInst(cpy3.PyInstanceMethod(fn.instance))
//	defer method.DecRef()
//	pClass.SetAttrString(fnName, method)
//	return PyClassFromObj(pClass)
//}
