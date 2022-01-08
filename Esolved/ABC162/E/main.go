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
	e := make([]int, k+1)
	for i := k; i >= 1; i-- {
		e[i] = Pow(k/i, n, p)
		for j := 2 * i; j <= k; j += i {
			e[i] -= e[j]
		}
	}
	//fmt.Println(n, e)
	ans := 0
	for i := 1; i <= k; i++ {
		ans += i * e[i]
		ans %= p
	}
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
