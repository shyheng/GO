package main

import "fmt"

func add() func(int) int {
	var n int = 10
	return func(i int) int {
		n += i
		return n
	}
}

func main() {
	f := add()
	fmt.Println(f(1))
}
