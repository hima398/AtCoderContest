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

	r, g, b := nextInt(), nextInt(), nextInt()
	ok := (100*r+10*g+b)%4 == 0
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
