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

	h, w := nextInt(), nextInt()
	ans := 0
	if h == 1 {
		ans = w
	} else if w == 1 {
		ans = h
	} else {
		//h >= 2 && w >= 2
		ans = Ceil(h, 2) * Ceil(w, 2)
	}
	fmt.Println(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
