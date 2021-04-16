package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	e := make([]int, n+1)
	r := 1
	max := 0
	for i := 2; i <= n; i++ {
		//fmt.Fprintf(out, "? %d %d\n", 1, i)
		fmt.Printf("? %d %d\n", 1, i)
		e[i] = nextInt()
		if max < e[i] {
			r = i
			max = e[i]
		}
	}
	//fmt.Println(r, max)
	ans := e[r]
	for i := 2; i <= n; i++ {
		if i == r {
			continue
		}
		fmt.Printf("? %d %d\n", r, i)
		d := nextInt()
		ans = Max(ans, d)
	}
	fmt.Fprintf(out, "! %d\n", ans)
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
