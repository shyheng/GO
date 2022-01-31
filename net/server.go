package main

import (
	"fmt"
	"net"
)

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

	}
}
