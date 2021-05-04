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

	n := nextInt()
	a := make([]int, n+1)
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		m[a[i]]++
	}
	//fmt.Println(m)
	s := 0
	c := make(map[int]int)
	for k, v := range m {
		c[k] = v * (v - 1) / 2
		s += c[k]
	}
	//fmt.Println(c)
	//fmt.Println(s)

	for i := 1; i <= n; i++ {
		if m[a[i]] < 1 {
			fmt.Println(s)
			continue
		}
		ans := s - c[a[i]] + (m[a[i]]-1)*(m[a[i]]-2)/2
		fmt.Println(ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
