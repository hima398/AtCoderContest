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

	n := nextInt()
	x0, y0 := nextInt(), nextInt()
	xn2, yn2 := nextInt(), nextInt()
	cx := float64(xn2+x0) / 2.0
	cy := float64(yn2+y0) / 2.0
	x := float64(x0) - cx
	y := float64(y0) - cy
	rad := 2.0 * math.Pi / float64(n)
	ansx := x*math.Cos(rad) - y*math.Sin(rad)
	ansy := x*math.Sin(rad) + y*math.Cos(rad)
	ansx += cx
	ansy += cy
	fmt.Println(ansx, ansy)

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
