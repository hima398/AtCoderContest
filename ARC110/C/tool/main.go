package main

import "fmt"

func main() {
	fmt.Println(200000)
	fmt.Printf("%d", 200000)
	for i := 1; i < 200000; i++ {
		fmt.Printf(" %d", i)
	}
}
