package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const P = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	x := make([]int, n+1)
	y := make([]int, n+1)
	z := make([]int, n+1)
	m := make([]int, n+1)
	for i := 0; i < n; i++ {
		x[i+1] = x[i]
		y[i+1] = y[i]
		z[i+1] = z[i]
		if m[a[i]] == 0 {
			x[i+1]++
		} else if m[a[i]] == 1 {
			y[i+1]++
		} else {
			z[i+1]++
		}
		m[a[i]]++
	}
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)
	//fmt.Println(m)
	ans := 1
	for i := 1; i <= n; i++ {
		ti := 0
		if x[i-1] == a[i-1] {
			ti++
		}
		if y[i-1] == a[i-1] {
			ti++
		}
		if z[i-1] == a[i-1] {
			ti++
		}
		ans *= ti
		ans %= P
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
