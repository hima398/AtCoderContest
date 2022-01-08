package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeDistance(sx, sy, tx, ty int) float64 {
	x := tx - sx
	y := ty - sy
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	txa, tya, txb, tyb, t, v := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	d := t * v
	for i := 0; i < n; i++ {
		dist := ComputeDistance(txa, tya, x[i], y[i]) + ComputeDistance(x[i], y[i], txb, tyb)
		if dist <= float64(d) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
