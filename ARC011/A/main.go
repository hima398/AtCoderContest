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

	m, n, ln := nextInt(), nextInt(), nextInt()
	i := ln
	d := 0
	ans := i
	for i/m > 0 {
		d = i % m
		i = (i / m) * n
		//fmt.Println(i)
		ans += i
		i += d
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
