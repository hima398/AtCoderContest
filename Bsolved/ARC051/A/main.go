package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Dist2(x1, y1, x2, y2 int) int {
	x, y := (x2 - x1), (y2 - y1)
	return x*x + y*y
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x1, y1, r := nextInt(), nextInt(), nextInt()
	x2, y2, x3, y3 := nextInt(), nextInt(), nextInt(), nextInt()
	// 青が赤に完全に含まれる
	r2 := r * r
	if r2 >= Dist2(x1, y1, x2, y2) && r2 >= Dist2(x1, y1, x2, y3) && r2 >= Dist2(x1, y1, x3, y2) && r2 >= Dist2(x1, y1, x3, y3) {
		fmt.Println("YES")
		fmt.Println("NO")
		return
	}
	// 赤が青に完全に含まれる
	if x2 <= x1 && x1 <= x3 && y2 <= y1 && y1 <= y3 && Abs(x2-x1) >= r && Abs(y2-y1) >= r && Abs(x3-x1) >= r && Abs(y3-y1) >= r {
		fmt.Println("NO")
		fmt.Println("YES")
		return
	}
	fmt.Println("YES")
	fmt.Println("YES")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
