package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	const p = 2019
	s := nextString()
	n := len(s)
	a := make([]int, n)
	for i := range a {
		a[i], _ = strconv.Atoi(string(s[i]))
	}
	t := make([]int, n+1)
	t[0] = 0
	j := 1
	for i := n - 1; i >= 0; i-- {
		t[n-i] = a[i]*j + t[n-i-1]
		t[n-i] %= p
		j *= 10
		j %= p
	}
	//fmt.Println(t)
	m := make(map[int]int)
	for _, v := range t {
		m[v]++
	}
	ans := 0
	for _, v := range m {
		if v > 1 {
			ans += v * (v - 1) / 2
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
