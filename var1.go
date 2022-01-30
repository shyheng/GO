package main

import "fmt"

func main() {
	a, b := 12, 50
	fmt.Println(a, b)
	a, b = b, a
	fmt.Print(a, b)
}
