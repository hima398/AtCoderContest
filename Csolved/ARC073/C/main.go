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

	ln, lt := nextInt(), nextInt()
	t := make([]int, ln+1)
	for i := 1; i <= ln; i++ {
		t[i] = nextInt()
	}
	ans := 0
	for i := 2; i <= ln; i++ {
		ans += Min(lt, t[i]-t[i-1])
	}
	ans += lt
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
