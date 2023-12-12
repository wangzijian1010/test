package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数 如何判断这个是什么模型
	// 1.方式1: 通过类型断言
	// 方法2 借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
}

func main() {
	a := 1.234
	reflectType(a)
}
