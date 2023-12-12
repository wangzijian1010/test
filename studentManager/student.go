package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

type studentMgr struct {
	allStudents []*student
}

// student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{id: id,
		name:  name,
		class: class}
}

// 管理系统的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		//如果新学生的id = 里面存在的id那么就是要编辑的id
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	// 如果走到最后还是没有的话 那么说明没有这个学生
	fmt.Println("输入学生信息有有误")
}

// 展示学生
func (s *studentMgr) showStudent() {
	// 这里可以将s.all当作迭代对象
	for _, v := range s.allStudents {
		fmt.Printf("学号%02d 姓名%s 班级%s \n", v.id, v.name, v.class)
	}
}
