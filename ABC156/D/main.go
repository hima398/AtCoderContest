package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const P = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, a, b int) int {
	ans := Pow(2, n, P) - 1
	//fmt.Println(ans)
	cna := ModCombination(n, a, P)
	for ans-cna <= 0 {
		ans += P
	}
	ans = (ans - cna) % P

	cnb := ModCombination(n, b, P)
	for ans-cnb <= 0 {
		ans += P
	}
	ans = (ans - cnb) % P
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a, b := nextInt(), nextInt(), nextInt()
	fmt.Println(Solve(n, a, b))
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y&1 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}

func ModFanctional(n, p int) int {
	if n == 1 {
		return 1
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans *= (n - i)
		ans %= p
	}
	return ans
}

func ModCombination(n, k, p int) int {
	x := 1
	mink := Min(k, n-k)
	for i := 0; i < mink; i++ {
		x *= (n - i)
		x %= p
	}
	y := ModFanctional(mink, p)
	return x * Inv(y, p) % p
}
