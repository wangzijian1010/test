package main

import "fmt"

// 为什么需要接口

type dog struct {
	name string
}

type cat struct {
	name string
}

type person struct {
	name string
}

func (d dog) Speak() {
	fmt.Println("汪汪汪")
}

func (c cat) Speak() {
	fmt.Println("喵喵喵")
}

func (p person) Speak() {
	fmt.Println("giao giao giao ")
}

// 其实就是为了解决如何一个方法被多个不同的对象调用的问题
// 人要吃东西 猪也要吃东西 吃东西这个行为是公用的 但是如果写太多不仅冗余 还容易出错
// 利用抽象的类型来代替不同的结构体 来进行抽象的方法
// 接口不管你是什么类型 他只关注你要实现什么方法

// 这里的正确理解就是 我i不管你是什么结构体dog也好cat也要 只要你有Speak这个方法 我就把你当作sayer这个类型
type sayer interface {
	Speak()
}

// 那这里就是sayer类型了 不管是dog还是cat都是sayer类型
func do(arg sayer) {
	arg.Speak()
}

func main() {
	c1 := cat{name: "cat"}
	do(c1)

	d1 := dog{name: "dog"}
	do(d1)

	p1 := person{name: "wangzijian"}
	do(p1)
}
