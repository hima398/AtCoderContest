package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SelectColor(r int) int {
	if r >= 1 && r < 400 {
		return 1
	} else if r < 800 {
		return 2
	} else if r < 1200 {
		return 3
	} else if r < 1600 {
		return 4
	} else if r < 2000 {
		return 5
	} else if r < 2400 {
		return 6
	} else if r < 2800 {
		return 7
	} else if r < 3200 {
		return 8
	} else {
		return 9
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := make(map[int]int)
	for i := 0; i < n; i++ {
		a := nextInt()
		c[SelectColor(a)]++
	}
	if c[9] == 0 {
		fmt.Println(len(c), len(c))
	} else if c[9] == n {
		fmt.Println(1, c[9])
	} else {
		fmt.Println(len(c)-1, len(c)+c[9]-1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
