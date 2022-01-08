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

	const INF = 1 << 60
	n, h := nextInt(), nextInt()
	a, b, c, d, e := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	ans := INF
	for i := 0; i <= n; i++ {
		j := sort.Search(n-i, func(j int) bool {
			return h+b*i+d*j-(n-i-j)*e > 0
		})
		ans = Min(ans, a*i+c*j)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
