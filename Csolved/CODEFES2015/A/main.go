package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k, m, r := nextInt(), nextInt(), nextInt(), nextInt()
	s := nextIntSlice(n - 1)
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	sum := 0
	for i := 0; i < k-1; i++ {
		sum += s[i]
	}
	//fmt.Println(sum + m)
	//fmt.Println(r * k)
	if sum+m < r*k {
		fmt.Println(-1)
		return
	}
	ng, ok := -1, m
	Check := func(x int) bool {
		ss := make([]int, n-1)
		copy(ss, s)
		ss = append(ss, x)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] > ss[j]
		})
		sum := 0
		for i := 0; i < k; i++ {
			sum += ss[i]
		}
		return sum >= r*k
	}
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if Check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
