package main

import "fmt"

// 使用值接收者 和使用指针接收者实现接口的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int
}

//func (p person) move() {
//	fmt.Printf("%s在跑", p.name)
//}

// 使用指针接收者
func (p *person) move() {
	fmt.Printf("%s再跑", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在说", p.name)
}

func main() {
	var m mover
	var s sayer
	p1 := &person{name: "xiaowangzi", age: 18}
	m = p1
	m.move()
	println()
	s = &person{name: "wangzijian", age: 19}
	s.say()

	var a animal // 定义一个animal类型的变量
	a = p1
	a.move()
	a.say()

}
