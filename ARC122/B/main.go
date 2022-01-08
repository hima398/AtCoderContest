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
	a := nextIntSlice(n)
	sort.Ints(a)
	var x float64
	if len(a)%2 == 0 {
		x = float64(a[n/2]+a[n/2-1]) / 4.0
	} else {
		x = float64(a[n/2]) / 2.0
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		fai := float64(a[i])
		ans += x + fai - math.Min(fai, 2*x)
	}
	ans /= float64(n)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
