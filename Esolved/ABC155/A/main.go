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
	/*
		if (a == b && b == c && a != c) || (a == b && b != c && a == c) || (a != b && b == c && a == c) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	*/
	if (a == b && a == c && b == c) || (a != b && a != c && b != c) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
