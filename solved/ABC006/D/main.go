package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	const INF = 3*int(1e4) + 1
	n := nextInt()
	a := nextIntSlice(n)
	ans := 0
	cards := make([]int, n+1)
	for i := 1; i <= n; i++ {
		cards[i] = INF
	}
	cards[0] = -1
	for i := 0; i < n; i++ {
		idx := sort.Search(n, func(j int) bool {
			return a[i] < cards[j]
		})
		if cards[idx] < INF {
			ans++
		}
		cards[idx] = a[i]
		//fmt.Println(i, a[i], idx)
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
