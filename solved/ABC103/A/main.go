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

	a := nextIntSlice(3)
	var ans []int
	ans = append(ans, Abs(a[0]-a[1])+Abs(a[1]-a[2]))
	ans = append(ans, Abs(a[0]-a[2])+Abs(a[2]-a[1]))
	ans = append(ans, Abs(a[1]-a[0])+Abs(a[0]-a[2]))
	ans = append(ans, Abs(a[1]-a[2])+Abs(a[2]-a[0]))
	ans = append(ans, Abs(a[2]-a[0])+Abs(a[0]-a[1]))
	ans = append(ans, Abs(a[2]-a[1])+Abs(a[1]-a[0]))
	sort.Ints(ans)
	fmt.Println(ans[0])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
