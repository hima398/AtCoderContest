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

	r, x, y := nextInt(), nextInt(), nextInt()
	dd := x*x + y*y
	d := math.Sqrt(float64(dd))
	var ans int
	if d == float64(r) {
		ans = 1
	} else if d <= float64(r+r) {
		ans = 2
	} else {
		ans = int(math.Ceil(d / float64(r)))
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
