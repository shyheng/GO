package main

import "fmt"

func main() {
	//ln自动换行
	fmt.Println("shy")
	//普通的不换行
	fmt.Print("zph")

	//%d 整数 ， %f 浮点数 ， %t 布尔类型 ， %s 字符串类型 ，%c 字符类型
	a := 1
	b := 1.26
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)

}
