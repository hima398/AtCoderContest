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

	e := make([]bool, int(1e6)+1)
	for i := 2 + 2; i <= int(1e6); i += 2 {
		e[i] = true
	}
	for i := 3; i <= int(1e6); i += 2 {
		for j := 2 * i; j <= int(1e6); j += i {
			e[j] = true
		}
	}
	x := nextInt()
	for i := x; i <= int(1e6); i++ {
		if !e[i] {
			fmt.Println(i)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
