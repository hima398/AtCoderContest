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

	n, m := nextInt(), nextInt()
	a := make([]int, n+1) // i-indexed
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	sa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sa[i] = (sa[i-1] + a[i]) % m
	}
	//fmt.Println(sa)
	sam := make(map[int]int)
	sam[0] = 1
	ans := 0
	for i := 1; i <= n; i++ {
		ans += sam[sa[i]]
		sam[sa[i]]++
	}
	//fmt.Println(sam)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
