package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(f [][]int) {
	for _, fi := range f {
		fmt.Println(fi)
	}
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	lx, ly, rx, ry := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		lx[i], ly[i], rx[i], ry[i] = nextInt(), nextInt(), nextInt(), nextInt()
		max = Max(max, rx[i])
		max = Max(max, ry[i])
	}
	max++
	f := make([][]int, max)
	for i := 0; i < max; i++ {
		f[i] = make([]int, max)
	}
	for i := 0; i < n; i++ {
		f[ly[i]][lx[i]]++
		f[ry[i]][rx[i]]++
		f[ly[i]][rx[i]]--
		f[ry[i]][lx[i]]--
	}
	//PrintField(f)
	for i := 0; i < max; i++ {
		for j := 0; j < max-1; j++ {
			f[j+1][i] += f[j][i]
		}
	}
	for j := 0; j < max; j++ {
		for i := 0; i < max-1; i++ {
			f[j][i+1] += f[j][i]
		}
	}
	//PrintField(f)
	ans := make([]int, n+1)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			ans[f[i][j]]++
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
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
