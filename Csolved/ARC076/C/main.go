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
	nn, mm := 1, 1
	for i := 1; i <= n; i++ {
		nn *= i
		nn %= p
	}
	for i := 1; i <= m; i++ {
		mm *= i
		mm %= p
	}
	ans := nn * mm % p
	if n == m {
		ans *= 2
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
