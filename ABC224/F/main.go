package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = 998244353
	s := nextString()
	var ss []int
	for _, v := range strings.Split(s, "") {
		iv, _ := strconv.Atoi(v)
		ss = append(ss, iv)
	}
	n := len(s)
	// pow2[n] = 2**n
	pow2 := make([]int, n)
	// pow10[n] = 10**n
	pow10 := make([]int, n)
	pow2[0] = 1
	pow10[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = 2 * pow2[i-1] % p
		pow10[i] = 10 * pow10[i-1] % p
	}
	w := make([]int, n)
	w[n-1] = pow2[n-1]
	for i := 1; i < n; i++ {
		w[n-i-1] = (w[n-i] + pow10[i]*pow2[n-i-1] - pow10[i-1]*pow2[n-i-1] + p) % p
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += w[i] * ss[i]
		ans %= p
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
