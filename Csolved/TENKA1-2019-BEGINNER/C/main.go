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
	s := nextString()

	nw, nb := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			nb++
		} else {
			nw++
		}
	}
	//自分を含む左にある黒の数
	sb := make([]int, n)
	if s[0] == '#' {
		sb[0] = 1
	}
	for i := 1; i < n; i++ {
		sb[i] = sb[i-1]
		if s[i] == '#' {
			sb[i]++
		}
	}

	//自分より右にある白の数
	sw := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		sw[i] = sw[i+1]
		if s[i+1] == '.' {
			sw[i]++
		}
	}

	//テスト出力
	//fmt.Println(sb)
	//fmt.Println(sw)

	ans := nw // 全てを黒にする場合の変更数
	for i := 0; i < n-1; i++ {
		//変更する数 = (自分を含む左にある黒) + (自分より右にある白)
		v := sb[i] + sw[i]
		ans = Min(ans, v)
	}
	ans = Min(ans, nb) //全てを白にする場合の変更数
	fmt.Println(ans)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
