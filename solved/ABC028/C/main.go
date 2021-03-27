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

	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = nextInt()
	}
	m := make(map[int]bool)
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				m[a[i]+a[j]+a[k]] = true
			}
		}
	}
	var ans []int
	for k := range m {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	fmt.Println(ans[len(ans)-3])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
