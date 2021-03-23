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

	n := nextInt()
	v := make([]int, n)
	c := make([]int, n)
	vc := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		c[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		vc[i] = v[i] - c[i]
	}
	sort.Ints(vc)
	ans := 0
	for i := n - 1; i >= 0 && vc[i] > 0; i-- {
		ans += vc[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
