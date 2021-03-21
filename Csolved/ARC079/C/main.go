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

	n, m := nextInt(), nextInt()
	mid := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		mid[a] = append(mid[a], b)
		mid[b] = append(mid[b], a)
	}
	for _, t1 := range mid[0] {
		for _, t2 := range mid[t1] {
			if t2 == n-1 {
				fmt.Println("POSSIBLE")
				return
			}
		}
	}
	fmt.Println("IMPOSSIBLE")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
