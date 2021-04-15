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

	a, b, c := nextInt(), nextInt(), nextInt()
	fa, fb, fc := float64(a), float64(b), float64(c)
	F := func(t float64) float64 {
		return fa*t + fb*math.Sin(fc*t*math.Pi)
	}
	l := 0.0
	h := 200.0
	for math.Abs(F(l)-100.0) > 1e-6 {
		m := (l + h) / 2.0
		if F(m) < 100.0 {
			l = m
		} else {
			h = m
		}
		//fmt.Println(l, m)
	}
	//fmt.Println(l, F(l))
	fmt.Println(l)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
