package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "aa"
	s[1] = "bb"
	s[2] = "cc"
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "dd")
	s = append(s, "ee", "ff")
	fmt.Println(s)

	ss := make([]string, len(s))
	copy(ss, s)
	fmt.Println(ss)

	fmt.Println(s[2:4])
	fmt.Println(s[:4])
	fmt.Println(s[2:])
}
