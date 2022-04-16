package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"reflect"
)

/*
@author RandySun
@create 2022-03-20-16:55
*/
// 声明服务端
type Client struct {
	conn net.Conn
}

// 构造方法
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// 会话连接的结构体
type Session1 struct {
	conn net.Conn
}

// 构造方法
func NewSession1(conn net.Conn) *Session1 {
	return &Session1{conn: conn}
}

// 向连接中去写数据
func (s *Session1) Write1(data []byte) error {
	// 定义写数据的格式
	// 4字节头部 + 可变体的长度
	buf := make([]byte, 4+len(data))
	// 写入头部，记录数据长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// 将整个数据，放到4后边
	copy(buf[4:], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// 从连接读数据
func (s *Session1) Read1() ([]byte, error) {
	// 读取头部记录的长度
	header := make([]byte, 4)
	// 按长度读取消息
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 定义RPC交互的数据结构
type RPCData1 struct {
	// 访问的函数
	Name string
	// 访问时的参数
	Args []interface{}
}

// 编码
func encode1(data RPCData1) ([]byte, error) {
	//得到字节数组的编码器
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)
	// 编码器对数据编码
	if err := bufEnc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码
func decode1(b []byte) (RPCData1, error) {
	buf := bytes.NewBuffer(b)
	// 得到字节数组解码器
	bufDec := gob.NewDecoder(buf)
	// 解码器对数据节码
	var data RPCData1
	if err := bufDec.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}

// 实现通用的RPC客户端
// 传入访问的函数名
// fPtr指向的是函数原型
//var select fun xx(User)
//cli.callRPC("selectUser",&select)
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	// 通过反射，获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	// 需要另一个函数，作用是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value {
		// 处理参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 连接
		cliSession := NewSession1(c.conn)
		// 编码数据
		reqRPC := RPCData1{Name: rpcName, Args: inArgs}
		b, err := encode1(reqRPC)
		if err != nil {
			panic(err)
		}
		// 写数据
		err = cliSession.Write1(b)
		if err != nil {
			panic(err)
		}
		// 服务端发过来返回值，此时应该读取和解析
		respBytes, err := cliSession.Read1()
		if err != nil {
			panic(err)
		}
		// 解码
		respRPC, err := decode1(respBytes)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行nil转换
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	// 完成原型到函数调用的内部转换
	// 参数1是reflect.Type
	// 参数2 f是函数类型，是对于参数1 fn函数的操作
	// fn是定义，f是具体操作
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值，过程
	fn.Set(v)
}

// 定义用户对象
type User1 struct {
	Name string
	Age  int
}
// 用于测试用户查询的方法
func queryUser1(uid int) (User1, error) {
	user := make(map[int]User1)
	// 假数据
	user[0] = User1{"zs", 20}
	user[1] = User1{"ls", 21}
	user[2] = User1{"ww", 22}
	// 模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User1{}, fmt.Errorf("%d err", uid)
}
func test()  {
	// 客户端获取连接
	addr := "127.0.0.1:8000"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("err")
	}
	// 创建客户端对象
	cli := NewClient(conn)
	// 需要声明函数原型
	var query func(int) (User1, error)
	cli.callRPC("queryUser", &query)
	// 得到查询结果
	u, err := query(1)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(u)
}
func main() {
	test()
}
