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

	a, b, c, k := nextInt(), nextInt(), nextInt(), nextInt()
	if k <= a {
		fmt.Println(k)
		return
	} else if k <= a+b {
		fmt.Println(a)
		return
	} else if k <= a+b+c {
		fmt.Println(a - (k - a - b))
		return
	} else {
		fmt.Println(a - c)
		return
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
