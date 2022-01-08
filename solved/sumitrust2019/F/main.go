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

	t := nextIntSlice(2)
	a := nextIntSlice(2)
	b := nextIntSlice(2)

	d1 := (a[0] - b[0]) * t[0]
	d2 := (a[1] - b[1]) * t[1]

	if d1 > 0 {
		d1 *= -1
		d2 *= -1
	}

	if d1+d2 < 0 {
		fmt.Println(0)
	} else if d1+d2 == 0 {
		fmt.Println("infinity")
	} else {
		s := -d1 / (d1 + d2)
		t := (-d1) % (d1 + d2)
		if t != 0 {
			fmt.Println(s*2 + 1)
		} else {
			fmt.Println(s * 2)
		}
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
