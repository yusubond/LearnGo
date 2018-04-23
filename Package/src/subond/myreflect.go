package subond

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func MyReflect() {
	a := new(MyStruct)
	a.name = "subond"
	// reflect.TypeOf()返回Type接口类型
	// so, Tpye既是方法的集合，也是一种数据类型
	// http://docscn.studygolang.com/pkg/reflect/#TypeOf
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(a).NumMethod())
	fmt.Println(reflect.TypeOf(a).Name())
	fmt.Println(reflect.TypeOf(a).Size())
}
