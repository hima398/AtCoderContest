package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type CD struct {
	i, cd int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	cds := make([]CD, n+1)
	for i := 1; i <= n; i++ {
		cds[i] = CD{i, i}
	}
	listen := make([]int, m)
	for i := 0; i < m; i++ {
		listen[i] = nextInt()
	}
	p := 0
	for i := 0; i < m; i++ {
		cds[p].i, cds[listen[i]].i = cds[listen[i]].i, 0
		p = listen[i]
	}
	sort.Slice(cds, func(i, j int) bool { return cds[i].i < cds[j].i })
	//fmt.Println(cds)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, cds[i].cd)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
