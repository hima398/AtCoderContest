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

	n, k := nextInt(), nextInt()
	sn := make([]int, n+1)
	for i := 0; i < k; i++ {
		d := nextInt()
		for j := 0; j < d; j++ {
			a := nextInt()
			sn[a]++
		}
	}
	//fmt.Println(sn)
	ans := 0
	for i := 1; i <= n; i++ {
		if sn[i] == 0 {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
