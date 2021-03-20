package IOC

import (
	"reflect"
	"todo-web/err"
)

//部分预制件
//pathValue 路径参数，自动转换
type Value struct {
	value reflect.Value
}

//Get send the value into given data pointer
func (pv *Value) Get(target interface{}) err.Exception {
	inValue := reflect.ValueOf(target)
	//一级指针
	rawValue := inValue.Elem()

	//试图赋值
	if rawValue.Type() == pv.value.Type() {
		rawValue.Set(pv.value)
		return err.NoExcetion
	}
	return err.UnSupportData(
		rawValue.Kind().String() +
			" not thr same with " +
			pv.value.String())
}

func setValue(data reflect.Value) reflect.Value {
	t := Value{value: data}
	return reflect.ValueOf(t)
}

type TenmplateData struct {
	data map[string]interface{}
}

func newTemplateMap() *TenmplateData {
	t := TenmplateData{}
	t.data = make(map[string]interface{})
	return &t
}
func (t *TenmplateData) Set(key string, data interface{}) {
	t.data[key] = data
}

type ConxtextSeter struct {
	data map[string]interface{}
}

func newConxtextSeter() *ConxtextSeter {
	t := ConxtextSeter{}
	t.data = make(map[string]interface{})
	return &t
}
func (t *ConxtextSeter) Set(key string, data interface{}) {
	t.data[key] = data
}
