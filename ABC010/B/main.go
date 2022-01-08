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

	// 1, 0, 1, 0, 1, 0, 1, 0, 1
	// 1, 0, 2, 1, 0, 2, 1, 0, 2
	// 6 + 1 or 6 + 3
	n := nextInt()
	a := nextIntSlice(n)
	ans := 0
	for i := 0; i < n; i++ {
		t := a[i] % 6
		if t == 1 || t == 3 {
			continue
		} else if t == 2 || t == 4 {
			ans++
		} else if t == 5 {
			ans += 2
		} else if t == 0 {
			//ai=0はあり得ないので6以上の6の倍数の場合
			ans += 3
		}

	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
