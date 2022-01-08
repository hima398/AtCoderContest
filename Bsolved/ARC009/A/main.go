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

	n := nextInt()
	a, b := make([]int, n), make([]int, n)
	var ans int
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		ans += a[i] * b[i]
	}
	ans = ans * 105 / 100
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
