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

	const p = int(1e9) + 7
	n := nextInt()
	t := nextIntSlice(n)
	sort.Ints(t)
	s := make([]int, n+1)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + t[i]
		m[t[i]]++
	}
	//fmt.Println(s)
	ans1 := 0
	for i := 0; i <= n; i++ {
		ans1 += s[i]
	}
	fmt.Println(ans1)
	ans2 := 1
	for _, v := range m {
		for i := v; i > 0; i-- {
			ans2 *= i
			ans2 %= p
		}
	}
	fmt.Println(ans2)
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
