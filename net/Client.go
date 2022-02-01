package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "192.168.1.101:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	//	客服端发送数据退出
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	write, err := dial.Write([]byte(readString))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(write)

}
