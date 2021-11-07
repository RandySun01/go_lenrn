package main

import "fmt"

func main() {
	a := []int{2,3,5,7,11,13}
	fmt.Printf("a = %p\n",a)
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))

	a = a[:0]
	fmt.Printf("a = %p\n",a)
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))

	a = a[:4]
	fmt.Printf("a = %p\n",a)
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))

	a = a[2:]
	fmt.Printf("a = %p\n",a)
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))

	b := 1
	fmt.Println(b)
	fmt.Printf("b=%v",&b)
	fmt.Println()
	b =2
	fmt.Println(b)
	fmt.Printf("b=%v",&b)


	c := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	fmt.Println(c) //[30 31 33 34 35 36 37]
	fmt.Println(c[:2]) //[30 31 33 34 35 36 37]
	fmt.Println(c[3:]) //[30 31 33 34 35 36 37]

	c = append(c[:2], c[3:]...)
	fmt.Println(c) //[30 31 33 34 35 36 37]

}
