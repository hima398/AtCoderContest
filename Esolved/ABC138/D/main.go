package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	xs := make([]int, n)
	for i := 0; i < q; i++ {
		p, x := nextInt(), nextInt()
		p--
		xs[p] += x
	}
	//fmt.Println(xs)
	ans := make([]int, n)
	var Dfs func(i, pre, x int)
	Dfs = func(i, pre, x int) {
		//fmt.Println(i, x)
		ans[i] += x + xs[i]
		for _, v := range e[i] {
			if v != pre {
				Dfs(v, i, x+xs[i])
			}
		}
	}
	Dfs(0, -1, 0)
	//fmt.Println(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
