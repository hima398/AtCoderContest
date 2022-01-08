package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrimeFactorization(x int) map[int]int {
	ret := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			ret[i]++
			x /= i
		}

		if x == 1 {
			break
		}
	}
	if x != 1 {
		ret[x]++
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7

	n := nextInt()
	a := nextIntSlice(n)
	//最小公倍数をmap[(素数)]素数をx乗で管理する
	pm := make(map[int]int)
	for i := 0; i < n; i++ {
		ps := PrimeFactorization(a[i])
		for k, v := range ps {
			pm[k] = Max(pm[k], v)
		}
	}
	//fmt.Println(pm)
	lcm := 1

	for k, v := range pm {
		lcm *= Pow(k, v, p)
		lcm %= p
	}

	ans := 0
	for i := 0; i < n; i++ {

		ans += lcm * Inv(a[i], p) % p
	}
	ans %= p
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
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
