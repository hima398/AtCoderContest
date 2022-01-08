package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(a, b int) (ans int) {

	ContainsInvalidNum := func(x int) bool {
		for x > 0 {
			if x%10 == 4 || x%10 == 9 {
				return true
			}
			x /= 10
		}
		return false
	}

	for i := a; i <= b; i++ {
		if ContainsInvalidNum(i) {
			ans++
		}
	}
	return ans
}

func Solve(a, b int) (ans int) {
	// x以下の数字で4、9を含まない数字を桁DBを使って計算して返す
	F := func(x int) int {
		sx := strconv.Itoa(x)
		n := len(sx)
		// xの左からi桁目まで見てi桁目の数字が小さければ0、同じであれば1に個数を記録する
		dp := make([][2]int, n+1)
		dp[0][0] = 1
		for i := 0; i < n; i++ {
			// i桁目の数字
			iv, _ := strconv.Atoi(string(sx[i]))
			for j := 0; j < 10; j++ {
				if j == 4 || j == 9 {
					continue
				}
				if j < iv {
					dp[i+1][1] += dp[i][0]
				}
				if j == iv {
					dp[i+1][0] += dp[i][0]
				}
				dp[i+1][1] += dp[i][1]
			}
			//fmt.Println(i, dp[i][0])
			//fmt.Println(i, dp[i][1])
		}
		ret := x - (dp[n][0] + dp[n][1])
		return ret
	}
	ans = F(b) - F(a-1)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	//ans := FirstSolve(a, b)
	ans := Solve(a, b)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
