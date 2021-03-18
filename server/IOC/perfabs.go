package IOC

import (
	"reflect"
	"todo-web/err"
)

//部分预制件
//pathValue 路径参数，自动转换
type Value struct {
	value interface{}
}

//Get send the value into given data pointer
func (pv *Value) Get(target interface{}) err.Exception {
	inValue := reflect.ValueOf(target)
	ownValue := reflect.ValueOf(pv.value)
	//一级指针
	rawValue := inValue.Elem()

	//试图赋值
	if rawValue.Type() == ownValue.Type() {
		rawValue.Set(ownValue)
		return err.NoExcetion
	}
	return err.UnSupportData(
		rawValue.Kind().String() +
			" not thr same with " +
			ownValue.Kind().String())
}
