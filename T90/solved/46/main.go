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

	const p = 46
	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)
	c := nextIntSlice(n)
	ma, mb, mc := make(map[int]int), make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		ma[a[i]%p]++
		mb[b[i]%p]++
		mc[c[i]%p]++
	}
	//fmt.Println(ma, mb, mc)
	ans := 0
	for i := 0; i < p; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < p; k++ {
				if (i+j+k)%46 == 0 {
					ans += ma[i] * mb[j] * mc[k]
				}
			}
		}
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
