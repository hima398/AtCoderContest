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
	x, y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	//a, b, c, d := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	r, b := make([]Pos, n), make([]Pos, n)

	for i := 0; i < n; i++ {
		r[i].x, r[i].y = nextInt(), nextInt()
	}
	for i := 0; i < n; i++ {
		b[i].x, b[i].y = nextInt(), nextInt()
	}
	sort.Slice(b, func(i, j int) bool {
		//if b[i].x < b[j].x {
		//	return true
		//}
		return b[i].y < b[j].y
	})
	//fmt.Println(b)
	sort.Slice(r, func(i, j int) bool {
		return r[i].x > r[j].x
	})

	ans := 0
	IsSelected := make([]bool, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !IsSelected[j] && r[i].x < b[j].x && r[i].y < b[j].y {
				IsSelected[j] = true
				ans++
				break
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
