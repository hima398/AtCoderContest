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

	_, m := nextInt(), nextInt()
	p := make([]int, m)
	y := make([]int, m)
	pref := make(map[int][]int)
	for i := 0; i < m; i++ {
		p[i], y[i] = nextInt(), nextInt()
		pref[p[i]] = append(pref[p[i]], y[i])
	}
	for k := range pref {
		sort.Ints(pref[k])
	}
	ans := make(map[int]map[int]int)
	for k := range pref {
		ans[k] = make(map[int]int)
		for i, v := range pref[k] {
			ans[k][v] = i + 1
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < m; i++ {
		fmt.Fprintf(out, "%06d%06d\n", p[i], ans[p[i]][y[i]])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
