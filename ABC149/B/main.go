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

	a, b, k := nextInt(), nextInt(), nextInt()
	if a < k {
		k -= a
		a = 0
	} else {
		a -= k
		fmt.Println(a, b)
		return
	}
	if b <= k {
		b = 0
	} else {
		b -= k
	}
	fmt.Println(a, b)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
