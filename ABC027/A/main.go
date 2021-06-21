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

	m := make(map[int]int)
	for i := 0; i < 3; i++ {
		l := nextInt()
		m[l]++
	}
	for k, v := range m {
		if v%2 == 1 {
			fmt.Println(k)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
