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

	y := nextInt()
	if y%400 == 0 {
		fmt.Println("YES")
		return
	}
	if y%100 == 0 {
		fmt.Println("NO")
		return
	}
	if y%4 == 0 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
