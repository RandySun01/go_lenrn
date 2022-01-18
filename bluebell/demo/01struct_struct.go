package main

import (
	"fmt"
	"unsafe"
)

/*
@author RandySun
@create 2022-01-18-8:50
*/

// 内存对齐
type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	c int8
	b string
}

func TestStruct() {
	v1 := s1{
		a: 1,
		b: "RandySun",
		c: 4,
	}
	v2 := s2{
		a: 1,
		c: 4,
		b: "RandySun",
	}
	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2))
}

func main() {
	TestStruct()
}
