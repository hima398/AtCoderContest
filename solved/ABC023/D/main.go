package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	h := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		h[i], s[i] = nextInt(), nextInt()
	}
	IsOk := func(x int) bool {
		t := make([]int, n)
		for i := 0; i < n; i++ {
			if x < h[i] {
				return false
			}
			t[i] = (x - h[i]) / s[i]
		}
		sort.Ints(t)
		//fmt.Println(x, t)
		for i, v := range t {
			if i > v {
				return false
			}
		}
		return true
	}
	ng, ok := 0, int(1e9)+int(1e5)*int(1e9)+1
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		//fmt.Printf("ng = %d, ok = %d, %b\n", ng, ok, IsOk(mid))
		if IsOk(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
