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

	const INF = 4*int(1e5) + 1
	n, m := nextInt(), nextInt()
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			d[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		a--
		b--
		d[a][b] = c
		d[b][a] = c
	}
	// O(n**3) ~ 12000000くらい
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				d[i][j] = Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	s := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			s += d[i][j]
		}
	}

	k := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for kk := 0; kk < k; kk++ {
		x, y, z := nextInt(), nextInt(), nextInt()
		x--
		y--

		//fmt.Fprintln(out, x, y, z, d[x][y])
		if d[x][y] > z {
			d[x][y] = z
			d[y][x] = z

			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					d[i][j] = Min(d[i][j], Min(d[i][x]+z+d[y][j], d[i][y]+z+d[x][j]))
				}
			}
			s = 0
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {
					s += d[i][j]
				}
			}
		}
		fmt.Fprintln(out, s)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
