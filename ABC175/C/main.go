package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x, k, d := nextInt(), nextInt(), nextInt()
	ax, ad := Abs(x), Abs(d)
	times := ax / ad
	if times >= k {
		fmt.Println(ax - d*k)
		return
	}
	ans := ax - d*times
	if (k-times)%2 == 0 {
		fmt.Println(ans)
		return
	} else {
		fmt.Println(Abs(Abs(ans) - ad))
		return
	}
}
