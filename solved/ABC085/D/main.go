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

	n, h := nextInt(), nextInt()
	a := make([]int, n)
	b := make([]int, n)
	maxa := 0
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		maxa = Max(maxa, a[i])
	}
	sort.Ints(b)
	ans := 0
	for i := n - 1; i >= 0 && (h > 0 && b[i] > maxa); i-- {
		h -= b[i]
		ans++
	}
	//fmt.Println(h, maxa)
	if h > 0 {
		if h%maxa == 0 {
			ans += h / maxa
		} else {
			ans += h/maxa + 1
		}
	}

	fmt.Println(ans)
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
