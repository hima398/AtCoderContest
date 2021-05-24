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
	a := nextIntSlice(n)
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			a[i] = -a[i]
		}
	}
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}
	//fmt.Println(s)

	m := make(map[int]int)
	ans := 0
	for i := 0; i <= n; i++ {
		//下記のコメントアウト部分と同等の計算
		ans += m[s[i]]
		m[s[i]]++
	}
	/*
		for i := 0; i <= n; i++ {
			m[s[i]]++
		}
		for _, v := range m {
			if v > 1 {
				ans += Combination(v, 2)
			}
		}
	*/

	fmt.Println(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}
