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

	n := nextInt()
	d := 0
	for i := n; i > 1; i /= 2 {
		d++
	}
	//fmt.Println(d)
	x := 1
	IsTakahashi := 1
	for x <= n {
		if IsTakahashi == 1 {
			if d%2 == 0 {
				x *= 2
				x++
			} else {
				x *= 2
			}
		} else {
			if d%2 == 0 {
				x *= 2
			} else {
				x *= 2
				x++
			}
		}
		IsTakahashi ^= 1
	}
	if IsTakahashi == 1 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
