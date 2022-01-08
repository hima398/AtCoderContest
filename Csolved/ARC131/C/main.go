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
	a := nextIntSlice(n)
	s := 0
	for i := 0; i < n; i++ {
		s ^= a[i]
	}

	if n%2 == 1 {
		fmt.Println("Win")
		return
	}
	for i := 0; i < n; i++ {
		if a[i] == s {
			fmt.Println("Win")
			return
		}
	}
	fmt.Println("Lose")
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
