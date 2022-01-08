package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const p = int(1e9) + 7

func SolveHonestly(h, w int) int {
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}
	dp[0][0] = 1

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ni, nj := i+1, j+1
			if ni < h {
				dp[ni][j] += dp[i][j]
				dp[ni][j] %= p
			}
			if nj < w {
				dp[i][nj] += dp[i][j]
				dp[i][nj] %= p
			}
		}
	}
	return dp[h-1][w-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	w, h := nextInt(), nextInt()
	//SolveHonestly(w, h)
	ans := ModCombination(h+w-2, h-1, p)
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

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}

func ModCombination(n, k, p int) int {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return n % p
	}
	ret := (n - k + 1) * ModCombination(n, k-1, p)
	ret %= p
	ret *= Inv(k, p)
	ret %= p
	return ret
}
