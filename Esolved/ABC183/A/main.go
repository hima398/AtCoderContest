package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	relu := x
	if x < 0 {
		relu = 0
	}
	fmt.Println(relu)
}
