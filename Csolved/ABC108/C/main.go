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

	n, k := nextInt(), nextInt()
	var v1 int
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			v1++
		}
	}
	ans := v1 * v1 * v1
	var v2 int
	if k%2 == 0 {
		for i := 1; i <= n; i++ {
			if i%k == k/2 {
				v2++
			}
		}
		ans += v2 * v2 * v2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
