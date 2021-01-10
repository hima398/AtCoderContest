package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n := nextInt()
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = nextInt()
	}
	sort.Ints(l)
	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				IsDiff := l[i] != l[j] && l[j] != l[k] && l[k] != l[i]
				IsTriangle := l[k] < l[i]+l[j]
				if IsDiff && IsTriangle {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
