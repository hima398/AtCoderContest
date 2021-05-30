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

	n, _ := nextInt(), nextInt()
	total := make(map[int]int)
	for i := 0; i < n; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			a := nextInt()
			total[a]++
		}
	}
	ans := 0
	for _, v := range total {
		if v == n {
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
