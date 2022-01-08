package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, x int, w []int) (ans int) {
	p1 := (1 << Min(n, 16)) - 1
	p2 := (1 << Max(n-16, 0)) - 1
	m1, m2 := make(map[int]int), make(map[int]int)

	//最大65535回ループ
	for pat := 0; pat <= p1; pat++ {
		ws := 0
		for i := 0; i < Min(n, 16); i++ {
			if (pat>>i)&1 == 1 {
				ws += w[i]
			}
		}
		m1[ws]++
	}
	//最大65535回ループ
	for pat := 0; pat <= p2; pat++ {
		ws := 0
		for i := 0; i < Max(n-16, 0); i++ {
			if (pat>>i)&1 == 1 {
				ws += w[i+16]
			}
		}
		m2[ws]++
	}
	//fmt.Println(m1, m2)
	//fmt.Println(p2, m2)
	ans = m1[x] + m2[x]

	for k := range m1 {
		k2 := x - k
		if k2 > 0 && k2 < x {
			ans += m1[k] * m2[k2]
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := nextInt(), nextInt()
	w := nextIntSlice(n)

	ans := Solve(n, x, w)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
