package main

import (
	"fmt"
	"strings"
)

//var wg sync.WaitGroup
//
//// 初始的例子
//
//func worker() {
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//	}
//	// 如何接收外部命令实现退出
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	// 如何优雅的实现结束子goroutine
//	wg.Wait()
//	fmt.Println("over")
//}
//var wg sync.WaitGroup
//var exit bool
//
//// 全局变量方式存在的问题
//
//// 1. 使用全局变量在跨包调用时不容易统一
//// 2. 如果worker中再启动goroutine，就不太好控制了。
//
//func worker() {
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		if exit {
//			break
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
//	exit = true                 // 修改全局变量实现子goroutine的退出
//	wg.Wait()
//	fmt.Println("over")
//}

//var wg sync.WaitGroup
//
//// 管道方式存在的问题：
//// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
//
//func worker(exitChan chan struct{}) {
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <-exitChan: // 等待接收上级通知
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	var exitChan = make(chan struct{})
//	wg.Add(1)
//	go worker(exitChan)
//	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
//	exitChan <- struct{}{}      // 给子goroutine发送退出信号
//	close(exitChan)
//	wg.Wait()
//	fmt.Println("over")
//}
//
//var wg sync.WaitGroup
//
//func worker(ctx context.Context) {
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <-ctx.Done(): // 等待上级通知
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 3)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}
//func gen(ctx context.Context) <-chan int {
//	dst := make(chan int)
//	n := 1
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return // return结束该goroutine，防止泄露
//			case dst <- n:
//				n++
//			}
//		}
//	}()
//	return dst
//}
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel() // 当我们取完需要的整数后调用cancel
//
//	for n := range gen(ctx) {
//		fmt.Println(n)
//		//if n == 5 {
//		//	break
//		//}
//	}
//}
//
//func main() {
//	d := time.Now().Add(50 * time.Millisecond)
//	ctx, cancel := context.WithDeadline(context.Background(), d)
//
//	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
//	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
//	defer cancel()
//
//	select {
//	case <-time.After(1 * time.Second):
//		fmt.Println("overslept")
//	case <-ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}

// context.WithTimeout

//var wg sync.WaitGroup
//
//func worker(ctx context.Context) {
//LOOP:
//	for {
//		fmt.Println("db connecting ...")
//		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
//		select {
//		case <-ctx.Done(): // 50毫秒后自动调用
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker done!")
//	wg.Done()
//}
//
//func main() {
//	// 设置一个50毫秒的超时
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}
//
//type TraceCode string
//
//var wg sync.WaitGroup
//
//func worker(ctx context.Context) {
//	key := TraceCode("TRACE_CODE")
//	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
//	if !ok {
//		fmt.Println("invalid trace code")
//	}
//LOOP:
//	for {
//		fmt.Printf("worker, trace code:%s\n", traceCode)
//		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
//		select {
//		case <-ctx.Done(): // 50毫秒后自动调用
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker done!")
//	wg.Done()
//}
//
//func main() {
//	// 设置一个50毫秒的超时
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
//	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
//	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")


//type Person struct {
//	Name  string
//	Hobby string
//}
//
//func main() {
//	p := Person{"5lmh.com", "女"}
//	// 编码json
//	b, err := json.Marshal(p)
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//
//	// 格式化输出
//	b, err = json.MarshalIndent(p, "", "     ")
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//
//}
//
//
//func main() {
//	student := make(map[string]interface{})
//	student["name"] = "5lmh.com"
//	student["age"] = 18
//	student["sex"] = "man"
//	b, err := json.Marshal(student)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(b))
//}


//type Person struct {
//Age       int    `json:"age,string"`
//Name      string `json:"name"`
//Niubility bool   `json:"niubility"`
//}
//
//func main() {
//	// 假数据
//	b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
//	var p Person
//	err := json.Unmarshal(b, &p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(p)
//}


//func main() {
//	// 假数据
//	// int和float64都当float64
//	b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)
//
//	// 声明接口
//	var i interface{}
//	err := json.Unmarshal(b, &i)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// 自动转到map
//	fmt.Println(i)
//	// 可以判断类型
//	m := i.(map[string]interface{})
//	for k, v := range m {
//		switch vv := v.(type) {
//		case float64:
//			fmt.Println(k, "是float64类型", vv)
//		case string:
//			fmt.Println(k, "是string类型", vv)
//		default:
//			fmt.Println("其他")
//		}
//	}
//}




//
//// 抽取单个server对象
//type Server struct {
//	ServerName string `xml:"serverName"`
//	ServerIP   string `xml:"serverIP"`
//}
//
//type Servers struct {
//	Name    xml.Name `xml:"servers"`
//	Version int   `xml:"version"`
//	Servers []Server `xml:"server"`
//}
//
//func main() {
//	data, err := ioutil.ReadFile("E:\\Lenrn_notes\\03前端\\04websocker\\websocker\\Go\\goWebsocket\\goroutine\\my.xml")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	var servers Servers
//	err = xml.Unmarshal(data, &servers)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("xml: %#v\n", servers)
//}


//
//type Person struct {
//	Name string
//	Age  int
//	Sex  string
//}
//
//// 二进制写出
//func writerJson(filename string) (err error) {
//	var persons []*Person
//	// 假数据
//	for i := 0; i < 10; i++ {
//		p := &Person{
//			Name: fmt.Sprintf("name%d", i),
//			Age:  rand.Intn(100),
//			Sex:  "male",
//		}
//		persons = append(persons, p)
//	}
//	// 二进制json序列化
//
//	data, err := msgpack.Marshal(persons)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = ioutil.WriteFile(filename, data, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	return
//}
//
//// 二进制读取
//func readJson(filename string) (err error) {
//	var persons []*Person
//	// 读文件
//	data, err := ioutil.ReadFile(filename)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 反序列化
//	err = msgpack.Unmarshal(data, &persons)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _, v := range persons {
//		fmt.Printf("%#v\n", v)
//	}
//	return
//}
//
//func main() {
//	err := writerJson("person.dat")
//	if err != nil {
//	fmt.Println(err)
//	return
//	}
//	//err := readJson("person.dat")
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//}

func main() {

	strSlice := strings.SplitAfter("a,b,c", ",")
	fmt.Println(strSlice,"\n")

	strSlice = strings.SplitAfter("I love my country", " ")
	for _, v := range strSlice {
		fmt.Println(v)
	}

	strSlice = strings.SplitAfter("abacadaeaf","a")
	fmt.Println("\n",strSlice)

	strSlice = strings.SplitAfter("abacadaeaf","A")
	fmt.Println("\n",strSlice)

	strSlice = strings.SplitAfter("123023403450456056706780789","0")
	fmt.Println("\n",strSlice)

	strSlice = strings.SplitAfter("123023403450456056706780789",",")
	fmt.Println("\n",strSlice)
}
