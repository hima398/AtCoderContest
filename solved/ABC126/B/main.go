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
	s := nextInt()
	f := s / 100
	r := s % 100
	if f >= 1 && f <= 12 && r >= 1 && r <= 12 {
		fmt.Println("AMBIGUOUS")
	} else if (f == 0 || f > 12) && r >= 1 && r <= 12 {
		fmt.Println("YYMM")
	} else if (r == 0 || r > 12) && f >= 1 && f <= 12 {
		fmt.Println("MMYY")
	} else {
		fmt.Println("NA")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
