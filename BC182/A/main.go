package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n := (2*a + 100) - b
	if n < 0 {
		n = 0
	}
	fmt.Println(n)
}
