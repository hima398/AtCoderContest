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
	m := make(map[string]int)
	max := 0
	for i := 0; i < n; i++ {
		s := nextString()
		m[s]++
		max = Max(max, m[s])
	}
	var ans []string
	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := range ans {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
