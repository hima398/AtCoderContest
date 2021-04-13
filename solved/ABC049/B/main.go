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

	h, _ := nextInt(), nextInt()
	c := make([]string, h)
	ans := make([]string, 2*h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
		ans[2*i] = c[i]
		ans[2*i+1] = c[i]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < 2*h; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
