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

	const INF = int(1e5) + 1
	n, c := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := INF
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				continue
			}
			cost := 0
			for k := 0; k < n; k++ {
				if k%2 == 0 {
					if a[k] != i {
						cost += c
					}
				} else {
					if a[k] != j {
						cost += c
					}
				}
			}
			ans = Min(ans, cost)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
