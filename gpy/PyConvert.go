package gpy

import (
	"reflect"
	"strings"
)

func convertCheckPyTypeName(toTypeString string,tp string)bool{
	if strings.HasPrefix(toTypeString,"*")&&strings.HasSuffix(toTypeString,tp){
		return true
	}
	return false
}

func GoToPyObject(o interface{}) IPyObject {
	to := reflect.TypeOf(o)
	valOf := reflect.ValueOf(o)

	if o == nil {
		return PyReturnNone()
	}
	//*PyObject return
	if to.AssignableTo(reflect.TypeOf(&PyObject{})) {
		return o.(*PyObject)
	}
	if valOf.CanInt() {
		return PyLongFromLongLong(valOf.Int())
	}
	if valOf.CanUint() {
		return PyLongFromLongLong(int64(valOf.Uint()))
	}
	if valOf.CanFloat() {
		return PyLong_FromDouble(valOf.Float())
	}

	if to.Kind() == reflect.Bool {
		return NewPyBool(valOf.Bool())
	}
	if to.Kind() == reflect.String {
		return NewPyUnicode(valOf.String())
	}

	if to.Kind() == reflect.Slice {
		l := NewPyList(0)
		size := valOf.Len()
		for i := 0; i < size; i++ {
			item := valOf.Index(i).Interface()
			o := GoToPyObject(item)
			defer o.DecRef()
			l.Append(o)
		}
		return l
	}
	if to.Kind() == reflect.Map {
		d := NewPyDict()
		keys := valOf.MapKeys()
		for _, key := range keys {
			k := GoToPyObject(key.Interface())
			defer k.DecRef()
			v := GoToPyObject(valOf.MapIndex(key).Interface())
			defer v.DecRef()
			d.SetItem(k, v)
		}
		return d
	}
	return PyReturnNone()
}

func PyObjectToGo(o IPyObject, to reflect.Type) any {
	//*PyObject return
	if to.AssignableTo(reflect.TypeOf(o)) {
		return o
	}
	toTypeString:=to.String()

	if convertCheckPyTypeName(toTypeString,"PyObject"){
		return AsPyObject(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyList"){
		return AsPyList(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyDict"){
		return AsPyDict(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyUnicode"){
		return AsPyUnicode(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyLong"){
		return AsPyLong(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyBool"){
		return AsPyBool(o)
	}
	if convertCheckPyTypeName(toTypeString,"PyBytes"){
		return AsPyBytes(o)
	}
	if convertCheckPyTypeName(toTypeString,"IPyObject"){
		return o
	}

	if to.Kind() == reflect.Int {
		return int(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Int8 {
		return int8(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Int16 {
		return int16(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Int32 {
		return int32(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Int64 {
		return int64(AsPyLong(o).AsInt64())
	}

	if to.Kind() == reflect.Uint {
		return uint(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Uint8 {
		return uint8(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Uint16 {
		return uint16(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Uint32 {
		return uint32(AsPyLong(o).AsInt64())
	}
	if to.Kind() == reflect.Uint64 {
		return uint64(AsPyLong(o).AsInt64())
	}

	if to.Kind() == reflect.Uintptr {
		return uintptr(AsPyLong(o).AsInt64())
	}

	if to.Kind() == reflect.Bool {
		return AsPyLong(o).AsInt() != 0
	}
	if to.Kind() == reflect.String {
		return o.ToString()
	}
	if to.Kind() == reflect.Float32 {
		return float32(AsPyLong(o).AsDouble())
	}
	if to.Kind() == reflect.Float64 {
		return float64(AsPyLong(o).AsDouble())
	}
	if to.Kind() == reflect.Slice {
		l := reflect.New(to).Elem()
		pyL := AsPyList(o)
		for i := int64(0); i < pyL.Size(); i++ {
			l = reflect.Append(l, reflect.ValueOf(PyObjectToGo(pyL.GetItem(i), to.Elem())))
		}
		return l.Interface()
	}
	if to.Kind() == reflect.Map {
		l := reflect.MakeMap(to)
		pyD := AsPyDict(o)
		if pyD.Size()>0{
			keys := AsPyList(pyD.Keys())
			defer keys.DecRef()
			for i := int64(0); i < keys.Size(); i++ {
				k := keys.GetItem(i)
				goK := PyObjectToGo(k, to.Key())
				l.SetMapIndex(reflect.ValueOf(goK), reflect.ValueOf(PyObjectToGo(pyD.GetItem(k), to.Elem())))
			}
		}
		return l.Interface()
	}
	if to.Kind() == reflect.Func {
		f := reflect.MakeFunc(to, func(args []reflect.Value) (results []reflect.Value) {
			pyArgs := NewPyTuple(int64(len(args)))
			for i, arg := range args {
				pyArgs.SetItem(int64(i), GoToPyObject(arg.Interface()))
			}
			rets := AsPyObject(o).Call(pyArgs, PyNil)
			//如果声明函数不需要返回
			if to.NumOut() == 0 {
				return nil
			}
			tp := rets.Type()
			defer tp.Free()

			goRets := make([]reflect.Value, 0)

			//如果是单个返回
			if tp.Name() != "tuple" {
				if to.NumOut() != 1 {
					k := make([]reflect.Value, 0)
					for i := 0; i < to.NumOut(); i++ {
						k = append(k, reflect.ValueOf(reflect.New(to.Out(i)).Elem().Interface()))
					}
					PyErr_SetString(PyExc_ValueError(), "返回参数不正确")
					return k
				}
				goRets = append(goRets, reflect.ValueOf(PyObjectToGo(rets, to.Out(0))))
			} else {
				tpRets := AsPyTuple(rets)
				if int64(to.NumOut()) != tpRets.Size() {
					k := make([]reflect.Value, 0)
					for i := 0; i < to.NumOut(); i++ {
						k = append(k, reflect.ValueOf(reflect.New(to.Out(i)).Elem().Interface()))
					}
					PyErr_SetString(PyExc_Exception(), "返回参数不正确")
					return k
				}
				for i := int64(0); i < tpRets.Size(); i++ {
					goRets = append(goRets, reflect.ValueOf(PyObjectToGo(tpRets.GetItem(i), to.Out(0))))
				}
			}
			return goRets
		})
		return f.Interface()
	}

	if to.Kind()==reflect.Interface{
		tp:=o.Type()
		defer tp.Free()
		if tp.ClassName()=="NoneType"{
			return nil
		}
		if tp.ClassName()=="str"{
			return PyObjectToGo(o,reflect.TypeOf(""))
		}
		if tp.ClassName()=="int"{
			return PyObjectToGo(o,reflect.TypeOf(int(1)))
		}
		if tp.ClassName()=="bool"{
			return PyObjectToGo(o,reflect.TypeOf(false))
		}
		if tp.ClassName()=="list"{
			return PyObjectToGo(o,reflect.TypeOf([]any{}))
		}
		if tp.ClassName()=="dict"{
			return PyObjectToGo(o,reflect.TypeOf(map[string]any{}))
		}
		if tp.ClassName()=="float"{
			return PyObjectToGo(o,reflect.TypeOf(float32(0)))
		}
	}


	return o
}
