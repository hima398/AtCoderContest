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
	if n == 1 {
		fmt.Println(1)
		return
	}
	a := nextIntSlice(n)
	var compress []int
	cnt := 1
	for i := 0; i < n-1; i++ {
		if a[i] < a[i+1] {
			cnt++
		} else {
			compress = append(compress, cnt)
			cnt = 1
		}
	}
	if cnt > 1 {
		compress = append(compress, cnt)
	} else {
		compress = append(compress, 1)
	}
	//fmt.Println(compress)
	ans := 0
	for _, v := range compress {
		ans += v * (v + 1) / 2
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
