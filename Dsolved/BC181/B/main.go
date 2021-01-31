package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
		sum += (b[i] + a[i]) * (b[i] - a[i] + 1) / 2
	}
	fmt.Printf("%d", sum)
}
