package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	p := make([][]int, h)
	for i := 0; i < h; i++ {
		p[i] = nextIntSlice(w)
	}

	max := (1 << h) - 1
	ans := 0
	for k := 1; k <= max; k++ {
		var r []int
		for j := 0; j < w; j++ {
			v := -1
			areSame := true
			for i := 0; i < h; i++ {
				if (k>>i)&1 == 0 {
					continue
				}
				if v < 0 {
					v = p[i][j]
				} else {
					areSame = areSame && v == p[i][j]
				}
			}
			if areSame {
				r = append(r, v)
			}
		}
		m := make(map[int]int)
		ph := bits.OnesCount(uint(k))
		pw := 0
		for _, v := range r {
			m[v]++
			pw = Max(pw, m[v])
		}
		ans = Max(ans, ph*pw)
	}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
