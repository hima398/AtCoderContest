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

	a, b, c, d := nextFloat64(), nextFloat64(), nextFloat64(), nextFloat64()
	if b/a > d/c {
		fmt.Println("TAKAHASHI")
	} else if b/a < d/c {
		fmt.Println("AOKI")
	} else {
		fmt.Println("DRAW")
	}
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
