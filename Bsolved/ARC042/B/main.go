package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeLine(x1, y1, x2, y2 int) (float64, float64, float64) {

	a := (y2 - y1)
	b := -(x2 - x1)
	c := -x1*y2 + x2*y1
	return float64(a), float64(b), float64(c)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1e5 + 1
	tx, ty := nextFloat64(), nextFloat64()
	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	ans := INF
	for i := 0; i < n; i++ {
		a, b, c := ComputeLine(x[i], y[i], x[(i+1)%n], y[(i+1)%n])
		d := math.Abs(a*tx + b*ty + c)
		d /= math.Sqrt(a*a + b*b)
		//fmt.Println(d, a, b, c, tx, ty)
		ans = math.Min(ans, d)
	}
	fmt.Println(ans)
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
