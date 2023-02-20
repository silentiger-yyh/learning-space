package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["two"])

	r, ok := m["un"]
	fmt.Println(r, ok)

	delete(m, "one")

	fmt.Println(m)

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"three": 3, "four": 4}
	fmt.Println(m2, m3)
}
