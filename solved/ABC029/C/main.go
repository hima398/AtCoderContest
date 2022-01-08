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

	pattern := []string{"a", "b", "c"}
	n := nextInt()
	var ans []string
	var F func(s string)
	F = func(s string) {
		if len(s) == n {
			ans = append(ans, s)
			return
		}
		for i := 0; i < len(pattern); i++ {
			F(s + pattern[i])
		}
	}
	for i := 0; i < len(pattern); i++ {
		F(pattern[i])
	}
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}
