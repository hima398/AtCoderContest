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

	const p = int(1e9) + 7

	n, m := nextInt(), nextInt()
	if Abs(n-m) > 1 {
		fmt.Println(0)
		return
	}
	if n < m {
		n, m = m, n
	}
	ans := 1
	if n-m == 1 {
		ans *= n
	} else {
		ans *= 2
	}
	for i := 0; i < m; i++ {
		ans *= m - i
		ans %= p
		ans *= m - i
		ans %= p
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
