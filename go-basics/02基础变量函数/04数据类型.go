package main

import "fmt"

/*

//变量类型
package main

import "fmt"

func main() {

*/
/*
	1 数字
		-int（整数，含正负）
			-int，int8，int16，int32，int64：占得大小不一样，
				-int在32位机器上是int32，64位机器上是int64
				-int8表示占一个字节（8个比特位）表示的范围是：正负2的7次方-1
				-int16表示占两个字节（16个比特位）表示的范围是：正负2的15次方-1
		-uint（正整数）
			-uint，uint8 uint16 uint32 uint64
				-uint在32位机器上是uint32，64位机器上是uint64
				-表示范围不一样uint8：2的8次方-1
		-float
			-float64和float32 ：表示小数点后范围不一样（，没有float）
		-complex64(复数：科学运算用得到，开发中，基本不用)了解
			-有实部和虚部
		-byte
			-uint8的别名
		-rune
			-int32的别名
	2 布尔(bool)
		-真：true
		-假：false
	3 字符串
		-双引号包裹的："xxx"
		-反引号包裹，可以换行，   没有单引号，三引号
*/
/*
//var a int=10
//var d float32=9.8
//var b bool=false
//var c bool=true
//var s string="lqz"
`` 是python中的三引号，可以换行
	var s string=`
你好
我好
大家好
`
var s string = "你好\n" +
	"我好" +
	"大家好"
fmt.Println(s)

}
*/
func main() {
	//var i int = 10
	//var s string = "str"
	//fmt.Println(string(i) + s)

	//	var sring string = `
	//
	//				1
	//2
	//	3
	//
	//				`
	var f float64
	fmt.Printf("%f\n", f)

	//fmt.Println(sring)
	//c := 123_456
	////print(c)
	//a := "sun"
	//b := "randy"
	//fmt.Println(a + b)
	////fmt.Sprintf(a, b)
	//demo := "I&love&Go,&and&I&also&love&Python."
	//string_slice := strings.Split(demo, "&")
	//s_d := strings.Split(a, "u")
	////print("s_d",s_d)
	//fmt.Println("s_d",s_d)
	//
	//fmt.Println("result:",string_slice)
	//fmt.Println("len:",len(string_slice))
	//fmt.Println("cap:", cap(string_slice))

	//fmt.Println(a)
	//print(len("s"))
	//fmt.Println(strings.Contains(a, "u"))

	//s := "hello沙河"
	//fmt.Println("len", len(s))
	//for i := 0; i < len(s); i++ { //byte
	//	fmt.Printf("%v(%c) ", s[i], s[i])
	//}
	//fmt.Println()
	//for a, r := range s { //rune
	//	fmt.Println(a,r)
	//	fmt.Printf("%v(%c) ", r, r)
	//	fmt.Printf("%c",  r)
	//
	//}
	//fmt.Println()

	//s1 := "big"
	//// 强制类型转换
	//byteS1 := []byte(s1)
	//byteS1[0] = 'p'
	//fmt.Println(string(byteS1))
	//
	//s2 := "白萝卜"
	//runeS2 := []rune(s2)
	//runeS2[0] = '红'
	//fmt.Println(string(runeS2))

	//s1 := "big"
	//// 强制类型转换
	//byteS1 := []byte(s1)
	//byteS1[0] = 'p'
	//fmt.Println(string(byteS1))
	//
	//s2 := "白萝卜"
	//runeS2 := []rune(s2)
	//runeS2[0] = '红'
	//fmt.Println(string(runeS2))
	//s3 := '你'
	//fmt.Println(s3)

	//s1:='汪' // 字符类型
	//fmt.Println(s1)
	//fmt.Printf("%T",s1) // int32,即rune类型
	//a :=1
	//a++
	//println(a)
	//
	//
	//a--
	//
	//println(a)
	//
	//var a = []int{1, 2, 3, 4, 5}
	//b := a    //此时a，b都指向了内存中的[1 2 3 4 5]的地址
	//b[1] = 10 //相当于修改同一个内存地址，所以a的值也会改变
	//
	//c := make([]int, 5, 5) //切片的初始化
	//copy(c, a)             //将切片a拷贝到c
	//c[1] = 20              //copy是拷贝值、创建了新的底层数组，所以a不会改变
	//
	//fmt.Printf("值a: %v，内存地址%p\n", a, &a) //a的值是[1 10 3 4 5]，a的内存地址是0xc42000a180
	//fmt.Printf("值b: %v，内存地址%p\n", b, &b) //b的值是[1 10 3 4 5]，b的内存地址是0xc42000a1a0
	//fmt.Printf("值c: %v，内存地址%p\n", c, &c) //c的值是[1 20 3 4 5]，c的内存地址是0xc42000a1c0
	//d := &a                           //将a的内存地址赋值给d，取值用*d
	//a[1] = 11
	//fmt.Printf("值是d: %v，内存地址%p\n", *d, d) //d的值是[1 11 3 4 5]，d的内存地址是0xc420084060
	//fmt.Printf("值是a: %v，内存地址%p\n", a, &a) //a的值是[1 11 3 4 5]，a的内存地址是0xc420084060
	//var a []string        //声明一个字符串切片，初始零值为？
	//fmt.Printf("%#v\n",a)
	//fmt.Println(a == nil) //true
	//
	//
	//var b = []int{}      //声明一个整型切片并初始，化初始零值为？
	//fmt.Printf("%#v\n",b)
	//fmt.Println(b == nil) //false

	//fmt.Println(nil== nil)
	//nil := "this is nil"
	//fmt.Println(nil)
	//nil := "this is nil"
	//fmt.Println(nil)
	//var slice []string = nil
	//fmt.Println(slice)

	//const val1 = iota
	//fmt.Printf("%T\n", val1)
	//var val2 = nil
	//fmt.Printf("%T\n", val2)

	// 指针类型的nil比较
	fmt.Println((*int64)(nil) == (*int64)(nil))
	// channel 类型的nil比较
	fmt.Println((chan int)(nil) == (chan int)(nil))
	// func类型的nil比较
	//fmt.Println((func())(nil) == (func())(nil)) // func() 只能与nil进行比较
	//// interface类型的nil比较
	fmt.Println((interface{})(nil) == (interface{})(nil))
	//// map类型的nil比较
	//fmt.Println((map[string]int)(nil) == (map[string]int)(nil)) // map 只能与nil进行比较
	//// slice类型的nil比较
	//fmt.Println(([]int)(nil) == ([]int)(nil)) // slice 只能与nil进行比较

	//var ptr *int64 = nil
	////var cha chan int64 = nil
	////var fun func() = nil
	//var inter interface{} = nil
	//var ma map[string]string = nil
	////var slice []int64 = nil
	////fmt.Println(ptr == cha)
	////fmt.Println(ptr == fun)
	//fmt.Println(ptr == inter)
	////fmt.Println(ptr == ma)
	////fmt.Println(ptr == slice)
	//fmt.Printf("%T\n", ma)
	//fmt.Printf("%T\n", inter)
	//fmt.Printf("%T\n", ptr)

	//var m map[string]string
	//fmt.Println(m["asoong"])
	//fmt.Println(m)
	//m["asong"] = "Golang梦工厂"

	var m *man
	fmt.Println(m.GetName())
}

type man struct {
	Name string
}

func (m *man) GetName() string {
	fmt.Println(m)
	return m.Name
}
