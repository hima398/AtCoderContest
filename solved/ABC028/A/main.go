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

	n := nextInt()
	if n <= 59 {
		fmt.Println("Bad")
	} else if n <= 89 {
		fmt.Println("Good")
	} else if n <= 99 {
		fmt.Println("Great")
	} else {
		fmt.Println("Perfect")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
