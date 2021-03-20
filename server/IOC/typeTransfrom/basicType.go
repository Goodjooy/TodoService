package typeTransfrom

import (
	"reflect"
	"strconv"
)

func toRaw(s interface{}) reflect.Value {
	return reflect.ValueOf(s)
}

func toString(s interface{}) reflect.Value {
	return reflect.ValueOf(s.(string))
}

func toInt(s interface{}) reflect.Value {
	i, e := strconv.ParseInt(s.(string), 10, 0)
	var r reflect.Value
	if e != nil {
		r = reflect.ValueOf(0)
	}
	r = reflect.ValueOf(i)

	return r
}
func toUint(s interface{}) reflect.Value {
	i, e := strconv.ParseUint(s.(string), 10, 0)
	var r reflect.Value
	if e != nil {
		r = reflect.ValueOf(0)
	}
	r = reflect.ValueOf(i)

	return r
}
