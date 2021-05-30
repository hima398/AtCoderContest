package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	match := map[int]int{1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6}
	n, m := nextInt(), nextInt()
	best := make(map[int]int)
	for i := 0; i < m; i++ {
		a := nextInt()
		best[match[a]] = Max(best[match[a]], a)
	}
	//fmt.Println(best)
	// i本のマッチ棒を使って作れる桁数の最大値
	dp := make([]string, n+1)
	for k, v := range best {
		if k <= n {
			dp[k] = strconv.Itoa(v)
		}
	}
	for i := 2; i <= n; i++ {
		for k, v := range best {
			pn := i - k
			if pn >= 0 && len(dp[pn]) > 0 {
				ndp := dp[pn] + strconv.Itoa(v)
				if len(dp[i]) < len(ndp) {
					dp[i] = ndp
				} else if len(dp[i]) == len(ndp) {
					if strings.Compare(dp[i], ndp) == -1 {
						dp[i] = ndp
					}
				}
			}
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[n])
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
