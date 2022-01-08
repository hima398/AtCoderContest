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

	n, lw := nextInt(), nextInt()
	w, v := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		w[i], v[i] = nextInt(), nextInt()
	}
	m := make([][]int, 4)
	w0 := w[0]
	for i := 0; i < n; i++ {
		idx := w[i] - w0
		m[idx] = append(m[idx], v[i])
	}
	for i := range m {
		sort.Slice(m[i], func(ii, jj int) bool {
			return m[i][ii] > m[i][jj]
		})
	}

	ans := 0
	for i := 0; i < len(m[0])+1; i++ {
		for j := 0; j < len(m[1])+1; j++ {
			for k := 0; k < len(m[2])+1; k++ {
				for l := 0; l < len(m[3])+1; l++ {
					tw := 0
					tv := 0
					for ii := 0; ii < i; ii++ {
						tw += w0
						tv += m[0][ii]
					}
					for jj := 0; jj < j; jj++ {
						tw += w0 + 1
						tv += m[1][jj]
					}
					for kk := 0; kk < k; kk++ {
						tw += w0 + 2
						tv += m[2][kk]
					}
					for ll := 0; ll < l; ll++ {
						tw += w0 + 3
						tv += m[3][ll]
					}
					if tw <= lw {
						//fmt.Println(i, j, k, l)
						ans = Max(ans, tv)
					}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
