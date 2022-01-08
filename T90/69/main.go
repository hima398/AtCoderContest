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
	n, k := nextInt(), nextInt()
	if k <= 2 && n > k {
		fmt.Println(0)
		return
	}
	if n == 1 {
		fmt.Println(k)
		return
	}

	ans := k * (k - 1)
	ans %= p
	// (k-2) ** n-2
	ans *= Pow(k-2, n-2, p)
	ans %= p
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}
