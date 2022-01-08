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
	a := nextIntSlice(n)
	ans := 0
	for k := 0; k <= 60; k++ {
		nz, no := 0, 0
		for i := 0; i < n; i++ {
			if a[i]&(1<<k) == 0 {
				nz++
			} else {
				no++
			}
		}
		//fmt.Println(k, nz, no)
		s := nz * no
		for j := 0; j < k; j++ {
			s *= 2
			s %= p
		}
		ans += s
		ans %= p
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
