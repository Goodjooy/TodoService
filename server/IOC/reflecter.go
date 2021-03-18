package IOC

import (
	"errors"
	"reflect"
	"todo-web/server/IOC/tools"
	"todo-web/server/manage"
)

//IOC 控制反转

var appMapping map[string][]FuncHandler

func addIOC(app manage.Application, handleFunc interface{}) error {
	valueType := reflect.TypeOf(handleFunc)
	if valueType.Kind() != reflect.Func {
		return errors.New("not a Func")
	}
	valueValue := reflect.ValueOf(handleFunc)

	var fun FuncHandler

	fun.fn = valueValue
	fun.fnType = valueType
	fun.inNum = uint(valueType.NumIn())
	fun.inArray = make([]InHandler, 0)
	ParmCount := fun.inNum

	//循环查找全部参数
	for i := 0; uint(i) < ParmCount; i++ {
		var inHandler InHandler
		parm := valueType.In(i)

		inHandler.parmType = parm
		inHandler.structType = parm.Kind() == reflect.Struct

		if inHandler.structType {
			//循环处理全部结构体变量
			for i := 0; i < parm.NumField(); i++ {
				var feild InFeildHandler
				f := parm.Field(i)
				feild.feildType = f.Type
				feild.name = f.Name
				feild.pkgPath = f.PkgPath
				feild.tag = f.Tag
				feild.targetType = tools.LoadTargetTypeTag(f.Tag)

				inHandler.insideFeild =
					append(inHandler.insideFeild, feild)
			}
		}

		fun.inArray = append(fun.inArray, inHandler)
	}

	return nil
}

func doIOC(appName string) {
	//todo generate new
	reflect.New(appMapping[appName][0].inArray[0].parmType)
}
