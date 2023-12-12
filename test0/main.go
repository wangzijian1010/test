package main

import (
	"encoding/json"
	"fmt"
)

// 如果go中标识是大写的 在其他包中才能访问 要不然只能在自己包里面可见
type student struct {
	ID   int
	Name string
}

// 想要Json序列化那你的变量名称就必须大写 因为是外部包想访问就得大写
// 但是也可以支持定制化 例如你的是json解析的时候 你就可以设置键值对的方法来进行 映射
// 其他的包也是类似的只要把json的名称改掉就行
type class struct {
	Title    string    `json:"title" db:"t"`
	Students []student `json:"students"`
}

// student的构造函数
// 主要就是给结构体赋值
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {
	// 创建一个班级变量c1
	// 切片必须要初始化
	c1 := class{Title: "火箭101",
		Students: make([]student, 0, 20),
	}
	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		// 新增10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v", c1)
	fmt.Println()

	fmt.Println()
	fmt.Println()

	// JSON序列化 Go语言中的数据转为Json字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("Json序列化失败")
		return
	}
	fmt.Printf("%s", data)
}
