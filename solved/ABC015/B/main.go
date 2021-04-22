package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	bn := 0
	s := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		bn += a
		if a > 0 {
			s++
		}
	}
	fmt.Println(math.Ceil(float64(bn) / float64(s)))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
