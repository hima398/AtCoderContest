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

	n, s, d := nextInt(), nextInt(), nextInt()
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		if x >= s || y <= d {
			continue
		}
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
