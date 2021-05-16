package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve() {
	const max = int(1e5)
	n := nextInt()
	e := make([]bool, max+1)
	e[0] = true
	e[1] = true
	for i := 2; i <= max; i++ {
		for j := i + i; j <= max; j += i {
			e[j] = true
		}
	}
	var p []int
	for i, IsNotPrime := range e {
		if !IsNotPrime {
			p = append(p, i)
		}
		if len(p) >= n {
			break
		}
	}
	//fmt.Println(len(p), p)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			// i != j
			ans[j] *= p[idx]
		}
		idx++
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out, "")

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	base := []int{6, 10, 15}
	ans := make([]int, 3)
	copy(ans, base)
	m := make(map[int]bool)
	for _, v := range base {
		m[v] = true
	}
	i := 16
	for len(ans) < n {
		if i%base[0] == 0 || i%base[1] == 0 || i%base[2] == 0 {
			ans = append(ans, i)
		}
		i++
	}
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
