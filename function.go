// function
package function

import (
	"fmt"
	"reflect"
)

type Function struct {
	handle     reflect.Value
	params     []reflect.Value
	paramIndex []uint8 // 回调参数序号数组
}

func (f *Function) IsValid() bool {
	return f.handle.IsValid()
}

func (f *Function) Call(p ...interface{}) []reflect.Value {
	var param []reflect.Value
	for i := 0; i < len(p); i++ {
		param = append(param, reflect.ValueOf(p[i]))
	}
	if len(f.paramIndex) != len(param) { //  如果参数出现不匹配情况那么直接报错 中止程序运行
		panic(fmt.Sprintf("参数不匹配 需要:%v 实际传入%v",
			len(f.paramIndex), len(param)))
	}
	in := make([]reflect.Value, len(f.params))
	copy(in, f.params)
	for i := 0; i < len(f.paramIndex); i++ {
		// 获取到参数索引位置
		index := in[f.paramIndex[i]].MethodByName("Get").Call([]reflect.Value{})
		rep := param[index[0].Uint()]
		if !rep.IsValid() { // 如果传递的参数不是合法的value  说明传递了 nil 尝试构造一个该参数的 0值
			in[f.paramIndex[i]] = reflect.Zero(f.handle.Type().In(int(f.paramIndex[i])))
		} else {
			in[f.paramIndex[i]] = rep
		}
	}
	return f.handle.Call(in)
}
