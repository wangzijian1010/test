package main

import "fmt"

// 空接口可以存放任何值
// 接口中没有定义任何方法时 为空接口
// 任意类型都实现了空接口 也就是空接口可以存放任意类型的值

// 空接口不需要提前定义一般 直接当临时存放的变量就行
type xxx interface {
}

func main() {
	// 任意类型都可以存储
	var x interface{}
	x = "hello"
	fmt.Println(x)

	x = 100
	fmt.Println(x)

	x = false
	fmt.Println(x)

}
