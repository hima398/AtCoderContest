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
	a := make([]int, n)
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		sorted[i] = a[i]
	}
	sort.Ints(sorted)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < n; i++ {
		if a[i] == sorted[n-1] {
			fmt.Println(sorted[n-2])
		} else {
			fmt.Println(sorted[n-1])
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
