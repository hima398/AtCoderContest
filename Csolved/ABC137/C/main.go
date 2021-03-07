package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s := strings.Split(nextString(), "")
		sort.Strings(s)
		m[strings.Join(s, "")]++
	}
	//fmt.Println(m)
	ans := 0
	for _, v := range m {
		if v <= 1 {
			continue
		}
		ans += Combination(v, 2)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Combination(N, K int) int {
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}
