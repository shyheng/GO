package main

import "fmt"

//枚举

func main()  {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	v := a
	fmt.Println(v)
	v = b
	fmt.Println(v)
}