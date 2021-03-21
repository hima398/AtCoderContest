package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	I int
	J int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, m := nextInt(), nextInt(), nextInt()
	row := make([]int, w)
	col := make([]int, h)
	b := make([]Pos, m)
	mrow := 0
	mcol := 0
	for i := 0; i < m; i++ {
		h, w := nextInt(), nextInt()
		h--
		w--
		b[i] = Pos{h, w}
		row[w]++
		col[h]++
		mrow = Max(mrow, row[w])
		mcol = Max(mcol, col[h])
	}
	var mj []int
	for j, v := range row {
		if v == mrow {
			mj = append(mj, j)
		}
	}
	var mi []int
	for i, v := range col {
		if v == mcol {
			mi = append(mi, i)
		}
	}
	ans := mcol + mrow - 1
	for _, iv := range mi {
		for _, jv := range mj {
			find := false
			//fmt.Println(iv, jv)
			for _, p := range b {
				find = find || (iv == p.I && jv == p.J)
			}
			if !find {
				ans++
				fmt.Println(ans)
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
