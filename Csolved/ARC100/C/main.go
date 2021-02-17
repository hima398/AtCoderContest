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
	//a := make([]int, n)
	an := make([]int, n)
	for i := 0; i < n; i++ {
		a := nextInt()
		an[i] = a - (i + 1)
	}
	//fmt.Println(an)
	sort.Ints(an)
	b := an[n/2]

	ans := 0
	for i := 0; i < n; i++ {
		ans += Abs(an[i] - b)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
