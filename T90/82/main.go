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

	const p = int(1e9) + 7
	l, r := nextInt(), nextInt()

	i2 := Inv(2, p)
	var F func(x int) int
	F = func(x int) int {
		if x == 0 {
			return 0
		}
		ret := 0
		v := 1
		for v <= x {
			s := (x % p) * ((x + 1) % p) % p
			s = s * i2 % p
			ret += s
			ret %= p
			t := (v % p) * ((v - 1) % p) % p
			t = t * i2 % p
			ret = ret - t + p
			ret %= p
			// int(1e18)を10倍してオーバーフローを検出する
			if v > v*10 {
				break
			}
			v *= 10
		}
		return ret % p
	}
	//fmt.Println(F(l-1), F(r))
	ans := F(r) - F(l-1) + p
	ans %= p
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
