package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < t; i++ {
		n, s, k := nextInt(), nextInt(), nextInt()
		g, x, _ := ExtGcd(k, n)
		if s%g != 0 {
			fmt.Fprintln(out, -1)
			continue
		}
		n /= g
		s /= g
		k /= g
		ans := (x * (-s)) % n
		if ans < 0 {
			ans += n
		}
		ans %= n
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	return Gcd(y, x%y)
}

// x, yの最大公約数gと
// a*x + b*y = gとなるa, bを返す
func ExtGcd(x, y int) (int, int, int) {
	if y == 0 {
		return x, 1, 0
	}
	g, a, b := ExtGcd(y, x%y)
	return g, b, a - x/y*b
}
