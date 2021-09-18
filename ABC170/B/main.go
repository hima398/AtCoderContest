package main

import (
	"bufio"
	"fmt"
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

	x, y := nextInt(), nextInt()
	for i := 0; i <= x; i++ {
		j := x - i
		if i*2+j*4 == y {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
