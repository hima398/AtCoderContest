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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, d := nextInt(), nextInt()
	ans := 0
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		dist := math.Sqrt(float64(x*x + y*y))
		if dist <= float64(d) {
			ans++
		}
	}
	fmt.Println(ans)
}
