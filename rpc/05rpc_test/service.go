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

// 定义RPC交互的数据结构
type RPCData struct {
	// 访问的函数
	Name string
	// 访问时的参数
	Args []interface{}
}

// 编码
func encode(data RPCData) ([]byte, error) {
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
func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	// 得到字节数组解码器
	bufDec := gob.NewDecoder(buf)
	// 解码器对数据节码
	var data RPCData
	if err := bufDec.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}

// 会话连接的结构体
type Session struct {
	conn net.Conn
}

// 构造方法
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中去写数据
func (s *Session) Write(data []byte) error {
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
func (s *Session) Read() ([]byte, error) {
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

// 声明服务端
type Server struct {
	// 地址
	addr string
	// map 用于维护关系的
	funcs map[string]reflect.Value
}

// 构造方法
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// 服务端需要一个注册Register
// 第一个参数函数名，第二个传入真正的函数
func (s *Server) Register(rpcName string, f interface{}) {
	// 维护一个map
	// 若map已经有键了
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// 若map中没值，则将映射加入map，用于调用
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

// 服务端等待调用的方法
func (s *Server) Run() {
	// 监听
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s err :%v", s.addr, err)
		return
	}
	for {
		// 服务端循环等待调用
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := NewSession(conn)
		// 使用RPC方式读取数据
		b, err := serSession.Read()
		if err != nil {
			return
		}
		// 数据解码
		rpcData, err := decode(b)
		if err != nil {
			return
		}
		// 根据读到的name，得到要调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Println("函数 %s 不存在", rpcData.Name)
			return
		}
		// 遍历解析客户端传来的参数,放切片里
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法
		// 返回Value类型，用于给客户端传递返回结果,out是所有的返回结果
		out := f.Call(inArgs)
		// 遍历out ，用于返回给客户端，存到一个切片里
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		// 数据编码，返回给客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respRPCData)
		if err != nil {
			return
		}
		// 将服务端编码后的数据，写出到客户端
		err = serSession.Write(bytes)
		if err != nil {
			return
		}
	}
}

//    给服务端注册一个查询用户的方法，客户端使用RPC方式调用

// 定义用户对象
type User struct {
	Name string
	Age  int
}

// 用于测试用户查询的方法
func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	// 假数据
	user[0] = User{"zs", 20}
	user[1] = User{"ls", 21}
	user[2] = User{"ww", 22}
	// 模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("%d err", uid)
}
func testService()  {
	// 编码中有一个字段是interface{}时，要注册一下
	//gob.Register(User{})
	gob.Register(User{})
	addr := "127.0.0.1:8000"
	// 创建服务端
	srv := NewServer(addr)
	// 将服务端方法，注册一下
	srv.Register("queryUser", queryUser)
	// 服务端等待调用
	srv.Run()
}
func main() {
	testService()

}
