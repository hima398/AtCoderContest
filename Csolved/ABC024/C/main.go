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

	_, d, k := nextInt(), nextInt(), nextInt()
	l, r := make([]int, d), make([]int, d)
	for i := 0; i < d; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}
	sl, sr, t := make([]int, k), make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		sl[i], t[i] = nextInt(), nextInt()
		sr[i] = sl[i]
	}

	ans := make([]int, k)
	for i := 0; i < d; i++ {
		for j := 0; j < k; j++ {
			if l[i] < sl[j] && sl[j] <= r[i] {
				sl[j] = l[i]
			}
			if sr[j] < r[i] && l[i] <= sr[j] {
				sr[j] = r[i]
			}
			if ans[j] == 0 && t[j] >= sl[j] && t[j] <= sr[j] {
				ans[j] = i + 1
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
