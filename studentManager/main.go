package main

import (
	"fmt"
	"os"
)

// 学员管理系统

// 1.添加学员
// 2.编辑学员信息
// 3.展示所有学员

func showMenu() {
	fmt.Println("欢迎来到学员管理系统")
	fmt.Println("1 添加学员")
	fmt.Println("2 编辑学员")
	fmt.Println("3 展示学员")
	fmt.Println("4 退出系统")
}

// 获取用户输入的学员信息
func getInput() *student {
	fmt.Println("请按照要求输入学员信息")
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("请输入学员的学号")
	fmt.Scan(&id)
	fmt.Print("请输入学员的姓名")
	fmt.Scan(&name)
	fmt.Print("请输入学员的班级")
	fmt.Scan(&class)
	// 返回的stu的结构体的指针
	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		// 打印系统菜单
		showMenu()

		// 等待用户选择选择要执行的选项
		fmt.Println("请输入想要使用的功能")
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("输入系统发生错误")
			return
		}

		switch number {
		case 1:
			// 实现添加学员的功能
			stu := getInput()
			sm.addStudent(stu)

		case 2:
			// 实现编辑学员的功能
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// 实现展示学员的功能
			sm.showStudent()
		case 4:
			// 退出功能
			os.Exit(0)

		}
	}

}
