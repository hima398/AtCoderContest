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
	s := make([]int, n)
	canAllDev := true
	sum := 0
	for i := 0; i < n; i++ {
		s[i] = nextInt()
		sum += s[i]
		canAllDev = canAllDev && s[i]%10 == 0
	}
	if canAllDev {
		fmt.Println(0)
		return
	}
	if sum%10 != 0 {
		fmt.Println(sum)
		return
	}
	sort.Ints(s)
	for _, v := range s {
		ans := sum - v
		if ans%10 != 0 {
			fmt.Println(ans)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
