package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly() {
	n, m := nextInt(), nextInt()
	e := make(map[int][]int)
	for i := 0; i < m; i++ {
		f, t := nextInt(), nextInt()
		f--
		t--
		e[f] = append(e[f], t)
		e[t] = append(e[t], f)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		m := map[int]bool{i: true}
		for _, v := range e[i] {
			m[v] = true
		}
		ans := make(map[int]int)
		for _, v := range e[i] {
			for _, v2 := range e[v] {
				if !m[v2] && ans[v2] <= 0 {
					ans[v2]++
				}
			}
		}
		fmt.Fprintln(out, len(ans))
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	SolveHonestly()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
