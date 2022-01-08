package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, v0, v1, l := nextInt(), nextFloat64(), nextFloat64(), nextFloat64()
	lt := 0.0
	for i := 0; i < n; i++ {
		//高橋くんが亀のところに行くまでの時間
		t := (l - lt) / v0
		lt = l
		l += v1 * t
	}
	fmt.Println(math.Max(0, l-lt))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
