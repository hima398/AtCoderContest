package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func F(x, p float64) float64 {
	return x + p/math.Pow(2, float64(x)*2.0/3.0)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	p := nextFloat64()

	l, r := 0.0, p
	//fl, fr := F(l, p), F(r, p)
	for (r-l)*1e5 >= 1.0 {
		//fmt.Println(l, fl, r, fr)
		c1 := l + (r-l)/3.0
		c2 := r - (r-l)/3.0

		if F(c1, p) < F(c2, p) {
			r = c2
			//fr = F(r, p)
		} else {
			l = c1
			//fl = F(l, p)
		}
	}
	ans := F(l, p)
	fmt.Println(ans)
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
