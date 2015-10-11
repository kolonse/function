// bind
package function

import (
	"reflect"
)

func Bind(handle interface{}, params ...interface{}) *Function {
	if reflect.TypeOf(handle).Kind() != reflect.Func {
		panic("传入参数不属于函数类型")
	}
	f := new(Function)
	f.handle = reflect.ValueOf(handle)
	globalPHType := reflect.TypeOf(globalPH)
	f.params = make([]reflect.Value, len(params))
	if params != nil && len(params) != 0 {
		for i := 0; i < len(params); i++ {
			f.params[i] = reflect.ValueOf(params[i])
			if reflect.TypeOf(params[i]) == globalPHType {
				f.paramIndex = append(f.paramIndex, uint8(i))
			}
		}
	}
	return f
}
