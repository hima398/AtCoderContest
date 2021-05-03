package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	if n%2 == 1 {
		return
	}
	memo := make([][]string, n+1)
	memo[2] = []string{"()"}
	// n = 2 or 4 or ... or 20
	var F func(n int) []string
	F = func(n int) []string {
		if n == 2 {
			return memo[2]
		}
		m := make(map[string]bool)
		for _, s := range F(n - 2) {
			m["("+s+")"] = true
		}
		for i := 2; i < n; i += 2 {
			for _, s1 := range F(i) {
				for _, s2 := range F(n - i) {
					m[s1+s2] = true
					m[s2+s1] = true
				}
			}
		}
		var ret []string
		for k := range m {
			ret = append(ret, k)
		}
		memo[n] = ret
		return memo[n]
	}
	ans := F(n)
	sort.Strings(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, s := range ans {
		fmt.Fprintln(out, s)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
