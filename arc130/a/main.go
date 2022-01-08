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
	s := nextString()
	var a []int
	cnt := 0
	for i := 0; i < n-1; i++ {
		cnt++
		if s[i] != s[i+1] {
			a = append(a, cnt)
			cnt = 0
		}
	}
	a = append(a, cnt+1)
	//fmt.Println(a)
	ans := 0
	for _, v := range a {
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
