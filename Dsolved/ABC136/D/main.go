package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

// RRR....LLLが与えられたとき、子供たちが集まる配列を返す
func F(s string) []int {
	ir, il := 0, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			ir = i
			il = i + 1
			break
		}
	}
	ret := make([]int, len(s))
	if len(s)%2 == 0 {
		ret[ir] = len(s) / 2
		ret[il] = len(s) / 2
	} else {
		if ir%2 == 0 {
			ret[ir] = len(s)/2 + 1
			ret[il] = len(s) / 2
		} else {
			ret[ir] = len(s) / 2
			ret[il] = len(s)/2 + 1
		}
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	var ans []int
	f, b := 0, 0
	for {
		if b == len(s)-1 || s[b] == 'L' && s[b+1] == 'R' {
			ans = append(ans, F(s[f:b+1])...)
			f = b + 1
			b = f
		} else {
			b++
		}
		if b >= len(s) {
			break
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", ans[0])
	for i := 1; i < len(s); i++ {
		fmt.Fprintf(out, " %d", ans[i])
	}
	fmt.Fprintln(out, "")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
