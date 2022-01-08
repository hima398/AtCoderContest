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

	n := nextInt()
	t, a := make([]int, n), make([]int, n)
	takahashi, aoki := 1, 1
	for i := 0; i < n; i++ {
		t[i], a[i] = nextInt(), nextInt()
	}
	for i := 0; i < n; i++ {
		x := Max(Ceil(takahashi, t[i]), Ceil(aoki, a[i]))
		takahashi = x * t[i]
		aoki = x * a[i]
	}
	fmt.Println(takahashi + aoki)
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

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
