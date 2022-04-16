package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

/*
@author RandySun
@create 2022-03-20-16:55
*/


// 测试网络中读写数据的情况

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

func main() {
	// 定义地址
	addr := "127.0.0.1:8000"
	my_data := "hello"
	// 等待组定义
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 写数据的协程
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		conn, _ := lis.Accept()
		s := Session{conn: conn}

		err = s.Write([]byte(my_data))

		if err != nil {
			log.Fatal(err)
		}
	}()

	// 读数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		s := Session{conn: conn}
		data, err := s.Read()
		if err != nil {
			log.Fatal(err)
		}
		// 最后一层校验
		if string(data) != my_data {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}