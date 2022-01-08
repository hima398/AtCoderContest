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
	m := map[int]int{3: 100000, 2: 200000, 1: 300000}
	x, y := nextInt(), nextInt()
	ans := m[x] + m[y]
	if x == 1 && y == 1 {
		ans += 400000
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
