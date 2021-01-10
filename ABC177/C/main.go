package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

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
	a := make([]int, n+1)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		s[i] = (s[i-1] + a[i]) % Mod
	}
	ans := 0
	for i := 1; i < n; i++ {
		ans += a[i] * (s[n] - s[i] + Mod) % Mod
	}
	fmt.Println(ans % Mod)
}
