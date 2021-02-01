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
	s := 0
	for i := 0; i < m; i++ {
		a := nextInt()
		s += a
	}
	if n < s {
		fmt.Println(-1)
		return
	}
	fmt.Println(n - s)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
