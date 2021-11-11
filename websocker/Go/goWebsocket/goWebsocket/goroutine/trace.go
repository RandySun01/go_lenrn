package main

import (
	"fmt"
	"time"
)

func main() {
	//
	////创建trace文件
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer f.Close()
	//
	////启动trace goroutine
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()
	//
	////main
	//fmt.Println("Hello World")

	//for i := 0; i < 5; i++ {
	//	time.Sleep(time.Second)
	//	fmt.Println("Hello World")
	//}

	//s1 := fmt.Sprint("枯藤")
	//name := "枯藤"
	//age := 18
	//s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	//s3 := fmt.Sprintln("枯藤")
	//fmt.Println(s1, s2, s3)

	//err := fmt.Errorf("这是一个错误")
	//fmt.Println(err)
	//
	//fmt.Printf("%v\n", 100)
	//fmt.Printf("%v\n", false)
	//o := struct{ name string }{"枯藤"}
	//fmt.Printf("%v\n", o)
	//fmt.Printf("%#v\n", o)
	//fmt.Printf("%T\n", o)
	//fmt.Printf("100%%\n")

	//f := 12.344444

	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scan(&name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	//reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	//fmt.Print("请输入内容：")
	//text, _ := reader.ReadString('\n') // 读到换行
	//
	//text = strings.TrimSpace(text)
	//fmt.Printf("%#v\n", text)

	//now := time.Now()            //获取当前时间
	//timestamp1 := now.Unix()     //时间戳
	//timestamp2 := now.UnixNano() //纳秒时间戳
	//fmt.Println(now)
	//fmt.Printf("current timestamp1:%v\n", timestamp1)
	//fmt.Printf("current timestamp2:%v\n", timestamp2)

	//timeObj := time.Unix(1630140365, 0) //将时间戳转为时间格式
	//fmt.Println(timeObj)
	//year := timeObj.Year()     //年
	//month := timeObj.Month()   //月
	//day := timeObj.Day()       //日
	//hour := timeObj.Hour()     //小时
	//minute := timeObj.Minute() //分钟
	//second := timeObj.Second() //秒
	//fmt.Printf("%d-%0?2d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	//now := time.Now()
	//later := now.Add(time.Hour) // 当前时间加1小时后的时间
	//fmt.Println(later)

	//ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	//for i := range ticker {
	//	fmt.Println(i)//每秒都会执行的任务
	//}
	//now := time.Now()
	//// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	//// 24小时制
	//fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//// 12小时制
	//fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	//fmt.Println(now.Format("2006/01/02 15:04"))
	//fmt.Println(now.Format("15:04 2006/01/02"))
	//fmt.Println(now.Format("2006/01/02"))
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
