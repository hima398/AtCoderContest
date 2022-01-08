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

	const p = int(1e9) + 7

	n := nextInt()
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	var F func(i, pre int) int
	var W func(i, pre int) int
	memoF := make([]int, n)
	memoW := make([]int, n)
	for i := 0; i < n; i++ {
		memoF[i] = -1
		memoW[i] = -1
	}
	F = func(i, pre int) int {
		if memoF[i] > 0 {
			return memoF[i]
		}
		if i == pre && len(e[i]) == 1 {
			memoF[i] = 2
			return memoF[i]
		}
		ret := 1
		for _, v := range e[i] {
			if v != pre {
				ret *= W(v, i)
				ret %= p
			}
		}
		ret += W(i, pre)
		ret %= p
		memoF[i] = ret
		return ret
	}
	W = func(i, pre int) int {
		if memoW[i] > 0 {
			return memoW[i]
		}
		if i == pre && len(e[i]) == 1 {
			memoW[i] = 1
			return memoW[i]
		}
		ret := 1
		for _, v := range e[i] {
			if v != pre {
				ret *= F(v, i)
				ret %= p
			}
		}
		memoW[i] = ret
		return memoW[i]
	}
	ans := F(0, -1)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
