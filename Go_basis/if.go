package main

import "fmt"

func main() {
	fmt.Println("请输入年份")
	var year int
	var b bool
	fmt.Scanf("%d",&year)
	b = (year%4 == 0) || (year%4 == 0&&year	%100 != 0)
	fmt.Printf("%t",b)
}
