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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	a := make([]int, k+1)
	a[1] = 7 % k
	for i := 1; i < k; i++ {
		if a[i] == 0 {
			fmt.Println(i)
			return
		}
		a[i+1] = (a[i]*10 + 7) % k
	}
	if a[k] == 0 {
		fmt.Println(k)
	} else {
		fmt.Println(-1)
	}
}
