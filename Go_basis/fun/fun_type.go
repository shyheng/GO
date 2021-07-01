package main

import "fmt"

func test()  {
	fmt.Println("你好")
}

func test1(a int,b int)  {
	fmt.Println(a+b)
}

type FU func()
type FU1 func(int,int)

func main() {
	var f FU
	f= test
	f()

	var f1 FU1
	f1 = test1
	f1(20,10)

}