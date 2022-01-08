package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FactorizePrime(n int) (ret map[int]int) {
	ret = make(map[int]int)
	rem := n
	for i := 2; i*i <= rem; i++ {
		for rem%i == 0 {
			ret[i]++
			rem /= i
		}
		if rem == 1 {
			break
		}
	}
	if rem != 1 {
		ret[rem]++
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	n, m := nextInt(), nextInt()
	ps := FactorizePrime(m)
	//fmt.Println(n, ps)

	ans := 1
	for _, v := range ps {
		ans *= ModCombination(n+v-1, n-1, p)
		ans %= p
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func ModCombination(n, k, p int) int {
	if k > n-k {
		return ModCombination(n, n-k, p)
	}
	mul, div := 1, 1
	for i := 0; i < k; i++ {
		mul *= n - i
		mul %= p
		div *= i + 1
		div %= p
	}
	ret := mul * Inv(div, p) % p
	return ret
}
