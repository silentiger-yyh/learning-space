package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}
	t := time.Now()
	switch {
	case t.Hour() < 10:
		fmt.Println(1)
	case t.Hour() < 12:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}
