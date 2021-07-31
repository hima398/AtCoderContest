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

	a, k := nextInt(), nextInt()
	if k == 0 {
		fmt.Println(2*int(1e12) - a)
		return
	}
	t := a
	ans := 0
	for t < 2*int(1e12) {
		ans++
		t += 1 + k*t
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
