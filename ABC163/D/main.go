package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const p = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	// m個選んだ時に作れる和の個数
	s := make([]int, n+2) // 0 - n+1
	for i := 1; i <= n+1; i++ {
		min := i * (i - 1) / 2
		max := i * (n + 1 - i + n) / 2
		//fmt.Println(min, max)
		s[i] = max - min + 1
	}
	//fmt.Println(k, s)
	ans := 0
	for i := k; i <= n+1; i++ {
		ans += s[i]
		ans %= p
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
