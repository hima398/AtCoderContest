package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FisrtSolve() {
	ln, ld, lh := nextInt(), nextInt(), nextInt()
	d := make([]int, ln)
	h := make([]int, lh)
	//md, mh := 0, 0
	for i := 0; i < ln; i++ {
		d[i], h[i] = nextInt(), nextInt()
		/*
			if mh < h[i] {
				mh = h[i]
				md = d[i]
			}
		*/
	}
	a := float64(lh) / float64(ld)
	for i := 0; i < ln; i++ {
		ai := float64(lh-h[i]) / float64(ld-d[i])
		/*
			if a > ai {
				a = ai
			}*/
		a = math.Min(a, ai)
	}
	ans := float64(lh) - a*float64(ld)
	fmt.Println(ans)
}

func Solve() {
	ln, ld, lh := nextInt(), nextInt(), nextInt()
	fd, fh := float64(ld), float64(lh)
	d := make([]float64, ln)
	h := make([]float64, ln)
	//md, mh := 0, 0
	ans := 0.0
	for i := 0; i < ln; i++ {
		d[i], h[i] = nextFloat64(), nextFloat64()
		a := (fh - h[i]) / (fd - d[i])
		y := fh - a*fd
		ans = math.Max(ans, y)
		/*
			if mh < h[i] {
				mh = h[i]
				md = d[i]
			}
		*/
	}
	fmt.Println(ans)

}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	FisrtSolve()
	//Solve()
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
