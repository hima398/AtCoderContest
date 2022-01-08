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

	a, b := nextInt(), nextInt()
	if a == 1 {
		a += 13
	}
	if b == 1 {
		b += 13
	}
	if a > b {
		fmt.Println("Alice")
	} else if a < b {
		fmt.Println("Bob")
	} else {
		fmt.Println("Draw")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
