package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if num := 1; num < 0 {
		fmt.Println(num)
	} else {
		fmt.Println(2)
	}
}
