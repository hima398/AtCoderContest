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
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = nextIntSlice(10)
	}
	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextIntSlice(11)
	}
	max := 1<<10 - 1
	ans := -(int(1e9) + 1)
	for b := 1; b <= max; b++ {
		bnf := 0 // 利益
		for i := 0; i < n; i++ {
			cnt := 0 //i番目のお店と同時に店を開ける数
			for j := 0; j < 10; j++ {
				if (b>>j)&1 == 1 && f[i][j] == 1 {
					cnt++
				}
			}
			bnf += p[i][cnt]
		}
		ans = Max(ans, bnf)
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
