package main

import "fmt"

// 空接口可以存放任何值
// 接口中没有定义任何方法时 为空接口
// 任意类型都实现了空接口 也就是空接口可以存放任意类型的值

// 空接口不需要提前定义一般 直接当临时存放的变量就行
type xxx interface {
}

// 1.空接口类型作为函数的参数就可以实现接收任何变量

// 2.空接口可以作为map的value 可以实现map的value的任意类型

func main() {
	// 任意类型都可以存储
	var x interface{}
	x = "hello"
	fmt.Println(x)

	x = 100
	fmt.Println(x)

	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "娜扎"
	m["age"] = 17
	m["hobby"] = []string{"篮球", "足球"}
	fmt.Println(m)

}
