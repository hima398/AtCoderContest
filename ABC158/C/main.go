package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly() {
	a, b := nextInt(), nextInt()
	for x := 10; x <= 1667; x++ {
		if x*8/100 == a && x/10 == b {
			fmt.Println(x)
			return
		}
	}
	fmt.Println(-1)
}

func Solve() {
	a, b := nextInt(), nextInt()
	xamin := 25 * a / 2
	xamax := (25*a + 24) / 2
	xbmin := 10 * b
	xbmax := 10 * (b + 1)
	if xbmin <= xamax && xamin < xbmax {
		fmt.Println(Max(xamin, xbmin))
	} else {
		fmt.Println(-1)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	SolveHonestly()
	//Solve()
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
