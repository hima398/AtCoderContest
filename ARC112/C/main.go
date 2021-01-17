package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Pow(x, y, p int) int {
	res := 1
	for y > 0 {
		if y&1 > 0 {
			res *= x
			res %= p
		}
		x = x * x % p
		y >>= 1
	}
	return res
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}

func Div(x, y, p int) int {
	return x * Inv(y, p) % p
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = 998244353
	h, w, k := nextInt(), nextInt(), nextInt()
	//g := make([][]int, w+1)
	//var g [200001][200001]byte
	g := make([][]byte, h+1)
	dp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		g[i] = make([]byte, w+1)
		dp[i] = make([]int, w+1)
	}
	for i := 0; i < k; i++ {
		hi, wi, ci := nextInt(), nextInt(), nextString()
		g[hi][wi] = []byte(ci)[0]
	}

	dp[1][1] = Pow(3, h*w-k, p)
	inv3 := Inv(3, p)
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if i > 1 {
				if g[i-1][j] == 'D' || g[i-1][j] == 'X' {
					dp[i][j] += dp[i-1][j]
				} else if g[i-1][j] == 0 {
					d := dp[i-1][j] * inv3 % p
					d = d * 2 % p
					dp[i][j] += d
				}
			}
			if j > 1 {
				if g[i][j-1] == 'R' || g[i][j-1] == 'X' {
					dp[i][j] += dp[i][j-1]
				} else if g[i][j-1] == 0 {
					d := dp[i][j-1] * inv3 % p
					d = d * 2 % p
					dp[i][j] += d
				}
			}
			dp[i][j] %= p
		}
	}
	/*
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				fmt.Printf("%d ", dp[i][j]%p)
			}
			fmt.Println("")
		}
	*/
	fmt.Println(dp[h][w])
}
