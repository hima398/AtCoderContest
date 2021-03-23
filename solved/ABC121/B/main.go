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

	n, m, c := nextInt(), nextInt(), nextInt()
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = nextInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		p := 0
		for j := 0; j < m; j++ {
			aij := nextInt()
			p += aij * b[j]
		}
		if p+c > 0 {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
