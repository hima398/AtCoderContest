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

	a, b, h, m := float64(nextInt()), float64(nextInt()), nextInt(), nextInt()
	rad := 2.0 * math.Pi * (float64(h)/12.0 + float64(m)/(60.0*12.0) - float64(m)/60.0)
	ans := a*a + b*b - 2.0*a*b*math.Cos(rad)

	fmt.Println(math.Sqrt(ans))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
