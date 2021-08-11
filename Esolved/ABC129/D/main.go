package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func SolveCommentary(h, w int) int {
	s := strings.Repeat("#", w+2)
	f := make([]string, h+2)
	f[0] = s
	for i := 1; i <= h; i++ {
		f[i] = "#" + nextString() + "#"
	}
	f[h+1] = s
	l := make([][]int, h+2)
	r := make([][]int, h+2)
	u := make([][]int, h+2)
	d := make([][]int, h+2)
	for i := 0; i <= h+1; i++ {
		l[i] = make([]int, w+2)
		r[i] = make([]int, w+2)
		u[i] = make([]int, w+2)
		d[i] = make([]int, w+2)
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if f[i][j] == '.' {
				r[i][j] = r[i][j-1] + 1
			}
		}
		for j := w; j >= 1; j-- {
			if f[i][j] == '.' {
				l[i][j] = l[i][j+1] + 1
			}
		}
	}
	for j := 1; j <= w; j++ {
		for i := 1; i <= h; i++ {
			if f[i][j] == '.' {
				u[i][j] = u[i-1][j] + 1
			}
		}
		for i := h; i >= 1; i-- {
			if f[i][j] == '.' {
				d[i][j] = d[i+1][j] + 1
			}
		}
	}
	ans := 0
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			lights := l[i][j] + r[i][j] + u[i][j] + d[i][j] - 3
			ans = Max(ans, lights)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	fmt.Println(SolveCommentary(h, w))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
