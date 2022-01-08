package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	a := make([][]string, h)
	for i := 0; i < h; i++ {
		a[i] = strings.Split(nextString(), "")
	}
	ans := make([][]string, h+2)
	for i := 0; i < h+2; i++ {
		ans[i] = make([]string, w+2)
		for j := 0; j < w+2; j++ {
			ans[i][j] = "#"
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[i+1][j+1] = a[i][j]
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < h+2; i++ {
		fmt.Fprintln(out, strings.Join(ans[i], ""))
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
