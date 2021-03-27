package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	i, j int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, d := nextInt(), nextInt(), nextInt()
	a := make([][]int, h)
	m := make(map[int]Pos)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
			m[a[i][j]] = Pos{i, j}
		}
	}
	//fmt.Println(m)
	s := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		s[i] = make([]int, h*w/d+1)
	}
	for i := 1; i <= d; i++ {
		p := m[i]
		j := 1
		for a[p.i][p.j]+d <= h*w {
			np := m[a[p.i][p.j]+d]
			s[i][j] = s[i][j-1] + Abs(np.i-p.i) + Abs(np.j-p.j)
			p = np
			j++
		}
	}
	s[0] = s[d]
	q := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		di := l % d
		var ans int
		if l%d == 0 {
			ans = s[di][r/d-1] - s[di][l/d-1]
		} else {
			ans = s[di][r/d] - s[di][l/d]
		}
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
