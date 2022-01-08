package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputePrimes(n int) []int {
	e := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		for j := 2 * i; j <= n; j += i {
			e[j] = true
		}
	}
	var ret []int
	for i := 2; i <= n; i++ {
		if !e[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const maxA = 55555
	n := nextInt()
	ps := ComputePrimes(maxA)
	var ans []int
	for _, p := range ps {
		if p%5 == 1 {
			ans = append(ans, p)
		}
		if len(ans) == n {
			break
		}
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
