package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	t := 0
	ans := 0
	for i := 0; i < n; i++ {
		d1, d2 := nextInt(), nextInt()
		if d1 == d2 {
			t++
		} else {
			ans = Max(ans, t)
			t = 0
		}
	}
	ans = Max(ans, t)
	if ans >= 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
