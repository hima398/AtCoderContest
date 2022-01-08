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

	_ = nextInt()
	a, b := nextInt(), nextInt()
	k := nextInt()
	m := make(map[int]int)
	for i := 0; i < k; i++ {
		pi := nextInt()
		m[pi]++
	}
	for pi, v := range m {
		if pi == a || pi == b {
			fmt.Println("NO")
			return
		}
		if v >= 2 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
