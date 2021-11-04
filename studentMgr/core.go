package main

import "fmt"

/*
@author RandySun
@create 2021-11-04-8:38
*/
// 学生管理系统

// 1. 添加学员
// 2. 编辑学员
// 3. 删除学员
// 4. 展示学员

func inputStudentInfo() *Student {
	var id int
	var name string
	var class string

	fmt.Printf("请输入学生编号:")
	fmt.Scanf("%d\n", &id)
	fmt.Printf("请输入学生姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入学生班级:")
	fmt.Scanf("%s\n", &class)

	return NewStudent(id, name, class)
}

func inputStudentId() int{
	var id int

	fmt.Printf("请输入学生编号:")
	fmt.Scanf("%d\n", &id)
	return id
}
func showMenu() {
	fmt.Println(" 1. 添加学员")
	fmt.Println(" 2. 编辑学员")
	fmt.Println(" 3. 删除学员")
	fmt.Println(" 4. 展示学员")

}

func Run() {
	var stuMgr = NewStudentMgr()
	for {
		fmt.Println("欢迎来到学生管理系统")
		showMenu()
		fmt.Printf("请输入操作编号：")
		var choice int
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			stu := inputStudentInfo()
			stuMgr.addStudent(stu)
		case 2:

			stuMgr.showStudent()
			stu := inputStudentInfo()
			stuMgr.modifyStudent(stu)
		case 3:

			stuMgr.showStudent()
			id := inputStudentId()
			stuMgr.deleteStudent(id)
		case 4:
			stuMgr.showStudent()
		default:
			fmt.Println("请选择已有功能")
		}
	}

}
