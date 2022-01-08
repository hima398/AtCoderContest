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

	w, h, x, y := nextInt(), nextInt(), nextInt(), nextInt()
	ans1 := float64(w*h) / 2.0
	ans2 := 0
	if x*2 == w && y*2 == h {
		ans2 = 1
	}
	fmt.Println(ans1, ans2)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
