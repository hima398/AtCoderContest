package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func InArea(a, b, c, d float64) bool {
	return math.Abs(a-c)+math.Abs(b-d) <= 3
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	r1, c1, r2, c2 := nextFloat(), nextFloat(), nextFloat(), nextFloat()
	if r1 == r2 && c1 == c2 {
		fmt.Println(0)
		return
	}
	// 同じ直線状 3の範囲
	if InArea(r1, c1, r2, c2) {
		fmt.Println(1)
		return
	}
	// 傾きが1 or -1
	k := (c2 - c1) / (r2 - r1)
	if k == 1 || k == -1 {
		fmt.Println(1)
		return
	}

	if (int64(c2-r2+c1+r1))%2 == 0 || int64(c2+r2+c1-r1)%2 == 0 {
		fmt.Println(2)
		return
	}

	if math.Abs(c2-r2-(c1-r1)) <= 3 || math.Abs(c2+r2-(c1+r1)) <= 3 {
		fmt.Println(2)
		return
	}

	fmt.Println(3)
}
