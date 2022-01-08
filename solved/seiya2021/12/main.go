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

	x := nextInt()
	if x < 600 {
		fmt.Println(8)
	} else if x < 800 {
		fmt.Println(7)
	} else if x < 1000 {
		fmt.Println(6)
	} else if x < 1200 {
		fmt.Println(5)
	} else if x < 1400 {
		fmt.Println(4)
	} else if x < 1600 {
		fmt.Println(3)
	} else if x < 1800 {
		fmt.Println(2)
	} else {
		fmt.Println(1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
