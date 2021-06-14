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

	n, m, k := nextInt(), nextInt(), nextInt()

	if n > m+k {
		fmt.Println(0)
		return
	}
	ans := ModCombination(n+m, n, p)
	if k < n {
		ans -= ModCombination(n+m, n-k-1, p)
		if ans < 0 {
			ans += p
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ModCombination(n, k, p int) int {
	if k == 0 {
		return 1
	}
	if k > n-k {
		return ModCombination(n, n-k, p)
	}
	mul, div := 1, 1
	for i := 0; i < k; i++ {
		mul *= n - i
		mul %= p
		div *= i + 1
		div %= p
	}
	ret := mul * Inv(div, p) % p
	return ret
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

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
