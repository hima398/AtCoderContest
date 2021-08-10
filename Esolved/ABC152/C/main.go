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
	mp := make([]int, n+1)
	mp[0] = 0
	mp[1] = nextInt()
	ans := 1
	for i := 2; i <= n; i++ {
		a := nextInt()
		if mp[i-1] >= a {
			ans++
		}
		mp[i] = Min(a, mp[i-1])
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
