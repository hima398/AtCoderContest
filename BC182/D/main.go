package main

import (
	"fmt"
	"sort"
)

func Max(nums []int) int {
	b := sort.IntSlice(nums)
	sort.Sort(b)
	return b[len(b)-1]
}

func main() {
	// 1 <= n <= 200000
	var n int
	fmt.Scan(&n)
	// -10 ^ 8 <= a[i] <= 10 ^8
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	lenA := len(a)
	p := []int{}
	pq := []int{}
	q := []int{}
	pi := 0
	for i := 0; i < lenA; i++ {
		pi += a[i]
		p = append(p, pi)
		pq = append(pq, pi)
		q = append(q, Max(pq))
	}
	robot := 0
	max := 0
	for i := 0; i < lenA; i++ {
		if max < robot+q[i] {
			max = robot + q[i]
		}
		robot += p[i]
	}
	fmt.Println(max)
}
