package main

import "fmt"

func main()  {
	time := 250551
	fmt.Println("天",time/60/60/24%360)
	fmt.Println("时",time/60/60%24)
	fmt.Println("分",time/60%60)
	fmt.Println("秒",time%60)
}
