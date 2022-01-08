package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	sort.Ints(x)
	sort.Ints(y)
	var mx, my float64
	if n%2 == 0 {
		mx = float64(x[n/2]) + float64(x[n/2-1])
		mx /= 2.0
		my = float64(y[n/2]) + float64(y[n/2-1])
		my /= 2.0
	} else {
		mx, my = float64(x[n/2]), float64(y[n/2])
	}
	//fmt.Println(mx, my)
	ax, ay := 0.0, 0.0
	for i := 0; i < n; i++ {
		fx, fy := float64(x[i]), float64(y[i])
		ax += math.Abs(fx - mx)
		ay += math.Abs(fy - my)
	}
	ans := int(ax + ay)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
