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
	IsTakahashi := 1
	for i := 0; i < k; i++ {
		if IsTakahashi == 1 {
			if a%2 == 1 {
				a--
			}
			b += a / 2
			a /= 2
		} else {
			if b%2 == 1 {
				b--
			}
			a += b / 2
			b /= 2
		}
		IsTakahashi ^= 1
	}
	fmt.Println(a, b)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
