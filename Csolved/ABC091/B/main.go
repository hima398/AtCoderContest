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

	m := make(map[string]int)
	nb := nextInt()
	for i := 0; i < nb; i++ {
		si := nextString()
		m[si]++
	}
	nr := nextInt()
	for i := 0; i < nr; i++ {
		ti := nextString()
		m[ti]--
	}
	ans := 0
	for _, v := range m {
		ans = Max(ans, v)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
