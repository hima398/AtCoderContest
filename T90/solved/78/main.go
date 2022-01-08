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

	n, m := nextInt(), nextInt()
	//e := make(map[int][]int)
	cnts := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		if a > b {
			cnts[a]++
		} else if a < b {
			cnts[b]++
		}
	}
	//fmt.Println(cnts)
	ans := 0
	for _, cnt := range cnts {
		if cnt == 1 {
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
