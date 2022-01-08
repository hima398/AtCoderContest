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

	a, b, c := nextInt(), nextInt(), nextInt()
	if a > b && a > c {
		if b > c {
			fmt.Println(1)
			fmt.Println(2)
			fmt.Println(3)
		} else {
			// c > b
			fmt.Println(1)
			fmt.Println(3)
			fmt.Println(2)
		}
	} else if b > a && b > c {
		if a > c {
			fmt.Println(2)
			fmt.Println(1)
			fmt.Println(3)
		} else {
			// c > a
			fmt.Println(3)
			fmt.Println(1)
			fmt.Println(2)
		}
	} else {
		// cが一番大きい
		if a > b {
			fmt.Println(2)
			fmt.Println(3)
			fmt.Println(1)

		} else {
			// b > a
			fmt.Println(3)
			fmt.Println(2)
			fmt.Println(1)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
