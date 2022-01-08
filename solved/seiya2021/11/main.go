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
	if x < 40 {
		fmt.Println(40 - x)
	} else if x < 70 {
		fmt.Println(70 - x)
	} else if x < 90 {
		fmt.Println(90 - x)
	} else {
		fmt.Println("expert")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
