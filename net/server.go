package main

import (
	"fmt"
	"io"
	_ "io"
	"net"
)

func process(conn net.Conn) {
	//	循环接收客服端发来的数据
	defer conn.Close()
	for {
		//	创建一个新的切片
		bytes := make([]byte, 1024)
		read, err := conn.Read(bytes)
		if err == io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println(string(bytes[:read]))
	}
}

func main() {
	fmt.Println("_____")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("错了")
		return
	}
	defer listen.Close()

	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("错了")
		} else {
			fmt.Println(accept)
		}
		//	准备协成为客服端服务
		go process(accept)
	}
}
