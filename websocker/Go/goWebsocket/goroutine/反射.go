package main

import (
	"fmt"
	"reflect"
)

//反射获取interface类型信息

//func reflect_type(a interface{}) {
//	t := reflect.TypeOf(a)
//	fmt.Println("类型是：", t)
//	// kind()可以获取具体类型
//	k := t.Kind()
//	fmt.Println(k)
//	switch k {
//	case reflect.Float64:
//		fmt.Printf("a is float64\n")
//	case reflect.String:
//		fmt.Println("string")
//	}
//}
//
//
////反射获取interface值信息
//
//func reflect_value(a interface{}) {
//	v := reflect.ValueOf(a)
//	fmt.Println(v)
//	k := v.Kind()
//	fmt.Println(k)
//	switch k {
//	case reflect.Float64:
//		fmt.Println("a是：", v.Float())
//	}
//}
////反射修改值
//func reflect_set_value(a interface{}) {
//	v := reflect.ValueOf(a)
//	k := v.Kind()
//	switch k {
//	case reflect.Float64:
//		// 反射修改值
//		v.SetFloat(6.9)
//		fmt.Println("a is ", v.Float())
//	case reflect.Ptr:
//		// Elem()获取地址指向的值
//		v.Elem().SetFloat(7.9)
//		fmt.Println("case:", v.Elem().Float())
//		// 地址
//		fmt.Println(v.Pointer())
//	}}
//func main() {
//	var x float64 = 3.4
//	//reflect_type(x)
//	//reflect_value(x)
//
//	// 反射认为下面是指针类型，不是float类型
//	reflect_set_value(&x)
//	fmt.Println("main:", x)
//}


//// 定义结构体
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//// 绑方法
//func (u User) Hello() {
//	fmt.Println("Hello")
//}
//
//// 传入interface{}
//func Poni(o interface{}) {
//	t := reflect.TypeOf(o)
//	fmt.Println("类型：", t)
//	fmt.Println("字符串类型：", t.Name())
//	// 获取值
//	v := reflect.ValueOf(o)
//	fmt.Println(v)
//	// 可以获取所有属性
//	// 获取结构体字段个数：t.NumField()
//	for i := 0; i < t.NumField(); i++ {
//		// 取每个字段
//		f := t.Field(i)
//		fmt.Printf("%s : %v", f.Name, f.Type)
//		// 获取字段的值信息
//		// Interface()：获取字段对应的值
//		val := v.Field(i).Interface()
//		fmt.Println("val :", val)
//	}
//	fmt.Println("=================方法====================")
//	for i := 0; i < t.NumMethod(); i++ {
//		m := t.Method(i)
//		fmt.Println(m.Name)
//		fmt.Println(m.Type)
//	}
//
//}
//
//func main() {
//	u := User{1, "zs", 20}
//	Poni(u)
//}


//// 定义结构体
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//// 匿名字段
//type Boy struct {
//	User
//	Addr string
//}
//
//func main() {
//	m := Boy{User{1, "zs", 20}, "bj"}
//	t := reflect.TypeOf(m)
//	fmt.Println(t)
//	// Anonymous：匿名
//	fmt.Printf("%#v\n", t.Field(0))
//	// 值信息
//	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
//}


//// 定义结构体
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//// 修改结构体值
//func SetValue(o interface{}) {
//	v := reflect.ValueOf(o)
//	// 获取指针指向的元素
//	v = v.Elem()
//	// 取字段
//	f := v.FieldByName("Name")
//	if f.Kind() == reflect.String {
//		f.SetString("kuteng")
//	}
//}
//
//func main() {
//	u := User{1, "5lmh.com", 20}
//	SetValue(&u)
//	fmt.Println(u)
//}
//

//// 定义结构体
//type User struct {
//	Id   int
//	Name string
//	Age  int
//}
//
//func (u User) Hello(name string) {
//	fmt.Println("Hello：", name)
//}
//
//func main() {
//	u := User{1, "5lmh.com", 20}
//	v := reflect.ValueOf(u)
//	// 获取方法
//	m := v.MethodByName("Hello")
//	// 构建一些参数
//	args := []reflect.Value{reflect.ValueOf("6666")}
//	// 没参数的情况下：var args2 []reflect.Value
//	// 调用方法，需要传入方法的参数
//	m.Call(args)
//}
//
//

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}

