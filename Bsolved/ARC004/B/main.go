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
	d := nextIntSlice(n)
	if n == 1 {
		fmt.Println(d[0])
		fmt.Println(d[0])
		return
	} else if n == 2 {
		fmt.Println(d[0] + d[1])
		fmt.Println(Abs(d[0] - d[1]))
		return
	}
	sd := make([]int, n)
	sd[0] = d[0]
	for i := 1; i < n; i++ {
		sd[i] = sd[i-1] + d[i]
	}
	max := sd[n-1]
	min := sd[n-1]
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {

			var ds []int
			ds = append(ds, sd[i])
			ds = append(ds, sd[j]-sd[i])
			ds = append(ds, sd[n-1]-sd[j])
			sort.Ints(ds)
			if ds[0]+ds[1] > ds[2] {
				min = 0
			} else {
				min = Min(min, ds[2]-ds[1]-ds[0])
			}

		}
	}
	fmt.Println(max)
	fmt.Println(min)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
