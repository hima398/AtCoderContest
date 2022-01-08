package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrimeFactorization(n int) map[int]int {
	ret := make(map[int]int)
	for n%2 == 0 {
		ret[2]++
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n == 1 {
			break
		}
		for n%i == 0 {
			ret[i]++
			n /= i
		}
	}
	if n != 1 {
		ret[n]++
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, p := nextInt(), nextInt()
	d := PrimeFactorization(p)
	//fmt.Println(d)
	ans := 1
	for k, v := range d {
		if v/n > 0 {
			ans *= Pow(k, v/n, int(1e12)+1)
		}
	}
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
