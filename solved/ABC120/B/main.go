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

	a, b, k := nextInt(), nextInt(), nextInt()
	n := Max(a, b)
	var ans []int
	//a, bが小さいので全列挙する
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			ans = append(ans, i)
		}
	}
	idx := len(ans) - k
	fmt.Println(ans[idx])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
