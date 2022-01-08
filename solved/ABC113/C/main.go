package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type City struct {
	i, p, y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_, m := nextInt(), nextInt()
	mc := make(map[int][]City)
	for i := 0; i < m; i++ {
		cs := City{i, nextInt(), nextInt()}
		mc[cs.p] = append(mc[cs.p], cs)
	}
	for _, v := range mc {
		sort.Slice(v, func(i, j int) bool {
			return v[i].y < v[j].y
		})
	}
	ans := make([]string, m)
	for _, v := range mc {
		for i, c := range v {
			ans[c.i] = fmt.Sprintf("%06d%06d", c.p, i+1)
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < m; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
