package main

import (
	"fmt"
	"sort"
)

func Max(a []int) int {
	b := sort.IntSlice(a)
	return b[len(b)-1]
}

func main() {
	// 1 <= n <= 100
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	maxA := Max(a)
	maxCount := 0
	result := 0
	for k := 2; k <= maxA; k++ {
		count := 0
		for i := 0; i < n; i++ {
			if a[i]%k == 0 {
				count++
			}
		}
		if maxCount <= count {
			maxCount = count
			result = k
		}
	}

	fmt.Println(result)
}
