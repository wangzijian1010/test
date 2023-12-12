package main

import "fmt"

// 数组有局限性

func Func1(Name string) (final string) {
	final = "hello" + Name
	return final
}

func Func2(a int, b int) (result int) {
	result = a + b
	return result
}

// 可变参数
func Func3(a ...int) {
	fmt.Println(a)
}

// 如果可变参数和固定参数同时存在那么需要将可变参数放在后面
func Func4(a int, b ...int) {

}

// 返回多个返回值
func Func5(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

func add(x, y int) (sum int) {
	return x + y
}

// 这里的op就是上面的add函数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func sub(x, y int) int {
	return x - y
}

func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 定义一个函数返回的也是一个函数
// 这里就是闭包的概念 函数加上外层变量的引用
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func a1() {

}

type person struct {
	name string
	city string
	age  int
}

// 实现Go中的构造函数
// 返回person结构体的指针开销会稍微小很多
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 方法的初始化
// 方法的定义其实就是在函数之前加了一个接收者的概念
func (p *person) Dream() {
	fmt.Printf("%s的梦想是学好go语言", p.name)
}

// 修改年龄
func (p *person) setAge(newAge int) {
	p.age = newAge
}

// 结构体的匿名字段
// 也就是不给结构体的变量取名字 这样估计适合临时的变量
// 匿名字段不能冲突 不能有两个string这样就分不清 3 28 38 45
type student struct {
	Person
	Address
}

type Person struct {
	Name   string
	Gender string
	Age    int
}

type Address struct {
	Province string
	City     string
}

// 结构体的继承
type Animal struct {
	name string
}

func (a *Animal) Move() {
	fmt.Printf("%s会动", a.name)
}

// 这里就相当于继承了Animal的类
type Dog struct {
	Feet    int
	*Animal // 匿名嵌套 而且嵌套指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪", d.name)
}

func main() {

	d1 := &Dog{Feet: 8, Animal: &Animal{name: "lele"}}
	d1.wang()
	println()
	// 可以继承Animal的方法Move
	d1.Move()
	println()

	stu := student{Address: Address{"GuangDong", "ShenZhen"}, Person: Person{Age: 20, Gender: "male", Name: "wangzijian"}}
	fmt.Printf("%#v", stu)

	// 将返回的person对象的指针赋值给了p
	// 类似构造函数
	p := newPerson("wangzijian", "shenzhen", 20)
	fmt.Println(p)
	// 修改年龄的方法
	p.setAge(59)
	fmt.Println(p)
	p.Dream()
	fmt.Println()
	r := a("wangzijian2")
	r() // 这个时候相当于执行了a中的匿名函数

	r1, r2 := calc2(2)
	println(r1(1))
	println(r2(2))

	result := calc(10, 20, add)
	result1 := calc1(10, 20, sub)
	println(result)
	println(result1)
	// 类似压栈的操作
	println("start....")
	defer println(1)
	defer println(2)
	defer println(3)
	println("end....")

	// 定义匿名函数
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

}
