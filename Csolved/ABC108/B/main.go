package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var x, y [4]int
	x[0], y[0], x[1], y[1] = nextInt(), nextInt(), nextInt(), nextInt()
	vx, vy := x[1]-x[0], y[1]-y[0]
	//90度回転
	vx, vy = -vy, vx
	x[2], y[2] = vx+x[1], vy+y[1]
	//90度回転
	vx, vy = -vy, vx
	x[3], y[3] = vx+x[2], vy+y[2]
	fmt.Println(x[2], y[2], x[3], y[3])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
