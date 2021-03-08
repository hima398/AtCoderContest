package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const maxn = 2001

var mc [maxn][maxn]int

func ModCombination(n, k, p int) int {
	if n < k {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k == 1 {
		return n
	}
	if mc[n][k] > 0 {
		return mc[n][k]
	}
	/*
		ret := ModCombination(n, k-1, p) * (n + 1 - k)
		ret %= p
		ret = ret * Inv(k, p)
		ret %= p
	*/
	c1, c2 := ModCombination(n-1, k, p), ModCombination(n-1, k-1, p)
	ret := (c1 + c2) % p
	mc[n][k] = ret
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	const p = int(1e9) + 7
	n, k := nextInt(), nextInt()
	ans := make([]int, k+1)
	for i := 1; i <= k; i++ {
		ans[i] = ModCombination(n-k+1, i, p) * ModCombination(k-1, i-1, p)
		ans[i] %= p
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= k; i++ {
		fmt.Println(ans[i])
	}
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

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
