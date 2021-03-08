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
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = nextInt()
	}
	sort.Ints(d)
	l, r := d[n/2-1], d[n/2]
	fmt.Println(r - l)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
