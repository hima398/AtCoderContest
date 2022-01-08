package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	d := make([][]int, n)
	s := make([][]int, n+1)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = nextInt()
		}
	}
	for i := 0; i <= n; i++ {
		s[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//fmt.Println(i, j, len(d), len(d[i]))
			s[i+1][j+1] = s[i][j+1] + s[i+1][j] - s[i][j] + d[i][j]
		}
	}
	q := nextInt()
	p := make([]int, q)
	for i := 0; i < q; i++ {
		p[i] = nextInt()
	}
	//PrintField(d)
	//fmt.Println()
	//PrintField(s)
	//最大で2501
	m := make([]int, n*n+1)
	//O(N**4) = 625000程度
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for ii := 0; ii <= i; ii++ {
				for jj := 0; jj <= j; jj++ {
					area := (i - ii) * (j - jj)
					deli := s[i][j] - s[ii][j] - s[i][jj] + s[ii][jj]
					m[area] = Max(m[area], deli)
				}
			}
		}
	}
	for k := 0; k < len(m)-1; k++ {
		m[k+1] = Max(m[k+1], m[k])
	}
	//fmt.Println(m)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for k := 0; k < q; k++ {
		fmt.Fprintln(out, m[p[k]])
	}

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
