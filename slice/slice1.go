package main

import "fmt"

func main() {
	var s = []int{100, 200, 300}
	s = append(s, 300)

	fmt.Println(s)
}
