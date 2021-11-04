package main

import "fmt"

/*
@author RandySun
@create 2021-11-03-8:00
*/
type Student struct {
	id    int
	name  string
	class string
}

// 学生信息初始化
func NewStudent(id int, name, class string) *Student {
	return &Student{id, name, class}
}

type StudentMgr struct {
	allStudent []*Student
}

// 初始化
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		allStudent: make([]*Student, 0, 10),
	}

}

// 添加学生
func (sMgr *StudentMgr) addStudent(student *Student) {
	// StudentMgr要定义指针接收者
	for _, stu := range sMgr.allStudent {
		if stu.id == student.id {
			fmt.Println("添加学员已经存在,请输入其他同学信息!")
			return
		}
	}
	sMgr.allStudent = append(sMgr.allStudent, student)
	fmt.Println("添加成功")
}

// 修改学生
func (sMgr *StudentMgr) modifyStudent(student *Student) {
	// StudentMgr要定义指针接收者
	for index, stu := range sMgr.allStudent {
		if stu.id == student.id {
			sMgr.allStudent[index] = student
			fmt.Println("修改成功")
			return
		}
	}
	fmt.Println("学生信息不存在,请添加学员信息")
}

// 删除学生
func (sMgr *StudentMgr) deleteStudent(id int) {
	if len(sMgr.allStudent) == 0 {
		fmt.Println("学员信息为空")
		return
	}
	// StudentMgr要定义指针接收者
	for index, stu := range sMgr.allStudent {
		if stu.id == id {
			endIndex := index + 1
			sMgr.allStudent = append(sMgr.allStudent[:index], sMgr.allStudent[endIndex:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("学生信息不存在,请添加学员信息")
}

// 展示学生信息
func (sMgr StudentMgr) showStudent() {
	if len(sMgr.allStudent) == 0 {
		fmt.Println("学员信息为空")
		return
	}
	for _, stu := range sMgr.allStudent {
		fmt.Printf("学生编号: %d 学生姓名: %s 学生班级: %s\n", stu.id, stu.name, stu.class)
	}
}
