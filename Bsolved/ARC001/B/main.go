package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Count(d int) int {

	ret := d / 10
	d = d % 10

	ret += d / 5
	d = d % 5

	ret += d
	return ret
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	d1 := Abs(a - b)
	ans1 := Count(d1)

	d2 := Abs(a - b)
	offset := (5 - d2%5)
	d2 += offset
	ans2 := Count(d2) + offset
	ans := Min(ans1, ans2)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
