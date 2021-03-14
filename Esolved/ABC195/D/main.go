package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Item struct {
	w, v int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, m, q := nextInt(), nextInt(), nextInt()
	item := make([]Item, n)
	x := make([]int, m)
	for i := 0; i < n; i++ {
		item[i] = Item{nextInt(), nextInt()}
	}
	sort.Slice(item, func(i, j int) bool { return item[i].v > item[j].v })
	for i := 0; i < m; i++ {
		x[i] = nextInt()
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for k := 0; k < q; k++ {
		l, r := nextInt(), nextInt()
		cx := make([]int, m)
		copy(cx, x)

		px := append(cx[0:l-1], cx[r:]...)
		if len(px) == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		sort.Ints(px)
		carry := make([]bool, len(px))
		ans := 0
		for i := 0; i < n; i++ {
			for j := 0; j < len(px); j++ {
				if item[i].w <= px[j] && !carry[j] {
					ans += item[i].v
					carry[j] = true
					break
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
