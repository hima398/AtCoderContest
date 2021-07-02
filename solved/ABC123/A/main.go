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

	a := nextIntSlice(5)
	k := nextInt()
	n := 0
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if a[j]-a[i] > k {
				n++
			}
		}
	}
	if n > 0 {
		fmt.Println(":(")
	} else {
		fmt.Println("Yay!")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
