package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Cell struct {
	m, d int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	md := [13]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var ms [13]int
	for i := 1; i <= 12; i++ {
		ms[i] = ms[i-1] + md[i]
	}
	//fmt.Println(ms)
	cal := make([]int, 367)
	for i := 1; i <= 366; i += 7 {
		cal[i-1] = 1 //土曜日
		cal[i] = 1   //日曜日
	}
	n := nextInt()
	c := make([]Cell, n)
	for i := 0; i < n; i++ {
		msd := strings.Split(nextString(), "/")
		m, _ := strconv.Atoi(msd[0])
		d, _ := strconv.Atoi(msd[1])
		c[i] = Cell{m, d}
	}
	sort.Slice(c, func(i, j int) bool {
		if c[i].m == c[j].m {
			return c[i].d < c[j].d
		} else {
			return c[i].m < c[j].m
		}
	})
	//msd = make([]string, n)
	for i := 0; i < n; i++ {
		idx := ms[c[i].m-1] + c[i].d
		for idx < 366 && cal[idx] > 0 {
			idx++
		}
		cal[idx] = 2
	}
	ans := 2
	cnt := 0
	for i := 1; i <= 366; i++ {
		if cal[i] > 0 {
			cnt++
		} else {
			ans = Max(ans, cnt)
			cnt = 0
		}
	}
	ans = Max(ans, cnt)
	fmt.Println(ans)
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
