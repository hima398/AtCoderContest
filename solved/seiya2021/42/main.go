package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_, _ = nextInt(), nextInt()
	n := nextInt()
	ch, cw := make(map[int]int), make(map[int]int)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		ch[a[i]]++
		cw[b[i]]++
	}
	var si, sj []int
	for k := range ch {
		si = append(si, k)
	}
	for k := range cw {
		sj = append(sj, k)
	}
	sort.Ints(si)
	sort.Ints(sj)
	mi, mj := make(map[int]int), make(map[int]int)
	for i := 0; i < len(si); i++ {
		mi[si[i]] = i + 1
	}
	for j := 0; j < len(sj); j++ {
		mj[sj[j]] = j + 1
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, mi[a[i]], mj[b[i]])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
