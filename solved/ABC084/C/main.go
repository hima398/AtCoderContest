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

	n := nextInt()
	c := make([]int, n)
	s := make([]int, n)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		c[i], s[i], f[i] = nextInt(), nextInt(), nextInt()
	}
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t := 0
		for j := i; j <= n-1; j++ {
			if t < s[j] {
				// i駅の始発待ち
				t = s[j]
			} else if t%f[j] == 0 {
				// すぐ次の電車に乗れるケース
			} else {
				t = t + f[j] - t%f[j]
			}
			t += c[j]
		}
		ans[i] = t
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
