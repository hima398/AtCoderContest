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

	a, b, x := nextInt(), nextInt(), nextInt()
	if a*a*b >= 2*x {
		h := 2.0 * float64(x) / (float64(a) * float64(b))
		rad := math.Atan(h / float64(b))
		deg := rad * 180.0 / math.Pi
		fmt.Println(90.0 - deg)
	} else {
		w := 2.0*float64(x) - float64(a)*float64(a)*float64(b)
		w /= float64(a) * float64(a)
		rad := math.Atan((float64(b) - w) / float64(a))
		deg := rad * 180.0 / math.Pi
		fmt.Println(deg)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
