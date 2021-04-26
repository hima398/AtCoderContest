package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	x, y, z := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		x[i], y[i], z[i] = nextInt(), nextInt(), nextInt()
	}
	//部分集合のパターン
	pss := make([]uint, 1<<n)
	for i := range pss {
		pss[i] = uint(i)
	}
	sort.Slice(pss, func(i, j int) bool { return bits.OnesCount(pss[i]) < bits.OnesCount(pss[j]) })
	dp := make([]int, 1<<n)
	dp[0] = 1
	for _, ps := range pss {
		c := bits.OnesCount(ps)
		ok := true
		for j := 0; j < m; j++ {
			if c > x[j] {
				continue
			}
			//部分集合内にあるyi以下の数
			cnt := bits.OnesCount(ps & ((1 << y[j]) - 1))
			if cnt > z[j] {
				ok = false
				break
			}
		}
		if ok {
			for j := 0; j < n; j++ {
				if ps>>j&1 == 1 {
					dp[ps] += dp[ps-(1<<j)]
				}
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
