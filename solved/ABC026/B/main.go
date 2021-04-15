package main

import (
	"bufio"
	"fmt"
	"math"
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
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}
	sort.Ints(r)
	red := 1
	ans := 0.0
	for i := n - 1; i >= 0; i-- {
		c := float64(r[i]*r[i]) * math.Pi
		if red == 1 {
			ans += c
		} else {
			ans -= c
		}
		red ^= 1
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
