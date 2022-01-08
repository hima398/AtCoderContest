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
		if b == 1 {
			fmt.Println("Draw")
		} else {
			fmt.Println("Alice")
		}
	} else {
		if b == 1 {
			fmt.Println("Bob")
		} else if a == b {
			fmt.Println("Draw")
		} else if a > b {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
