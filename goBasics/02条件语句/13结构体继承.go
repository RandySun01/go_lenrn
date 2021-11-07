package main

import (
	"encoding/json"
	"fmt"
)

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

// 结构体标签
//Student 学生
type Student2 struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	// 结构体继承
	//d1 := &Dog{
	//	Feet: 4,
	//	Animal: &Animal{ //注意嵌套的是结构体指针
	//		name: "乐乐",
	//	},
	//}
	//d1.wang() //乐乐会汪汪汪~
	//d1.move() //乐乐会动！

	// 序列化
	//c := &Class{
	//	Title:    "101",
	//	Students: make([]*Student, 0, 200),
	//}
	////
	//for i := 0; i < 10; i++ {
	//	stu := &Student{
	//		Name:   fmt.Sprintf("stu%02d", i),
	//		Gender: "男",
	//		ID:     i,
	//	}
	//	c.Students = append(c.Students, stu)
	//}
	////
	////	fmt.Println(c)
	////	//JSON序列化：结构体-->JSON格式的字符串
	//	data,err := json.Marshal(c)
	//	fmt.Println("err", err)
	////
	//	if err!=nil{
	//		fmt.Println("json marshal failed")
	//		return
	//	}
	//	fmt.Println(data)
	//	fmt.Printf("json:%s\n", data)
	//
	//	c1 := &Class{}
	//	//JSON反序列化：JSON格式的字符串-->;结构体
	//	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	//	err = json.Unmarshal([]byte(str), c1)
	//	if err != nil {
	//		fmt.Println("json unmarshal failed!")
	//		return
	//	}
	//	fmt.Printf("%#v\n", c1)
	// 结构体标签（Tag）
		s1 := Student2{
			ID:     1,
			Gender: "男",
			name:   "沙河娜扎",
		}
		data, err := json.Marshal(s1)
		if err != nil {
			fmt.Println("json marshal failed!")
			return
		}

		fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"} // json str:{"id":1,"Gender":"男"}

}
