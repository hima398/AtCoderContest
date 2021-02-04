package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func SloveHonestly() int {
	n := nextInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	ans := 0
	for i := 0; i < n; i++ {
		cantDivide := true
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			cantDivide = cantDivide && a[i]%a[j] != 0
		}
		if cantDivide {
			ans++
		}
	}
	return ans
}

func Solve() int {
	n := nextInt()
	a := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		max = Max(max, a[i])
	}

	d := make([]int, max+1)

	sort.Ints(a)
	for _, v := range a {
		d[v]++
		if d[v] > 1 {
			continue
		}
		for i := 2 * v; i <= max; i += v {
			d[i]++
		}
	}
	//fmt.Println(d)
	ans := 0
	for _, v := range a {
		if d[v] == 1 {
			ans++
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//fmt.Println(SloveHonestly())
	fmt.Println(Solve())
}
