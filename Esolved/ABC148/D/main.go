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
	//a :=make([]int, n)
	ans := 0
	i := 1
	for j := 0; j < n; j++ {
		a := nextInt()
		if a == i {
			i++
		} else {
			ans++
		}
	}
	if i == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
