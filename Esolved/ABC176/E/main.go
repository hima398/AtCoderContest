package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	I int
	J int
}

type Row struct {
	i, n int
}
type Col struct {
	j, n int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, m := nextInt(), nextInt(), nextInt()
	row := make([]Row, h)
	col := make([]Col, w)
	for i := 0; i < h; i++ {
		row[i].i = i
	}
	for j := 0; j < w; j++ {
		col[j].j = j
	}
	t := make(map[int]map[int]bool)
	for i := 0; i < m; i++ {
		hi, wi := nextInt(), nextInt()
		hi--
		wi--
		if t[hi] == nil {
			t[hi] = make(map[int]bool)
		}
		t[hi][wi] = true
		row[hi].n++
		col[wi].n++
	}
	//fmt.Println(t)
	sort.Slice(row, func(i, j int) bool {
		return row[i].n > row[j].n
	})
	sort.Slice(col, func(i, j int) bool {
		return col[i].n > col[j].n
	})
	mr, mc := row[0].n, col[0].n
	ans := mr + mc - 1
	for i := 0; i < h && row[i].n == mr; i++ {
		for j := 0; j < w && col[j].n == mc; j++ {
			if !t[row[i].i][col[j].j] {
				ans++
				fmt.Println(ans)
				return
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
