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

	const maxn = int(1e7)
	e := make([]bool, maxn+1)
	e[0], e[1] = true, true
	for i := 2; i <= maxn; i++ {
		for j := i + i; j <= maxn; j += i {
			e[j] = true
		}
	}
	p := make(map[int]bool)
	for i, v := range e {
		if !v {
			p[i] = true
		}
	}
	//fmt.Println(p)
	f := make([]int, maxn+1)
	for k := range p {
		for i := k; i <= maxn; i += k {
			f[i]++
		}
	}

	n, k := nextInt(), nextInt()
	ans := 0
	for i := 2; i <= n; i++ {
		if f[i] >= k {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
