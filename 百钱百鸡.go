package main

import "fmt"

func main() {

	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken++ {
				if cock+hen+chicken == 100 && cock*5+hen*3+chicken/3 == 100 {
					fmt.Print("公鸡", cock)
					fmt.Print("母鸡", hen)
					fmt.Print("小鸡", chicken)
					fmt.Println()
				}
			}
		}
	}
}
