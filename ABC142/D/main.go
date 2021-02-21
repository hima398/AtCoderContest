package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const max = int(1e12)
const rootMax = int(1e6)

var soe [rootMax + 1]bool

func CommonPrime(x, y int) []int {
	if x == 1 || y == 1 {
		return []int{}
	}
	m := make(map[int]bool)
	for i := 2; i <= rootMax; i++ {
		if x%i == 0 && y%i == 0 {
			if !soe[i] {
				m[i] = true
			}
		}
	}
	var ret []int
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	// a, bが互いに素
	if Gcd(a, b) == 1 {
		fmt.Println(1)
		return
	}
	for i := 2; i <= rootMax; i++ {
		if soe[i] {
			continue
		}
		for j := 2 * i; j <= rootMax; j += i {
			soe[j] = true
		}
	}
	d := CommonPrime(a, b)
	//fmt.Println(d)
	//fmt.Println(soe)
	ans := 1 + len(d)

	fmt.Println(ans)
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
	if x < y {
		x, y = y, x
	}
	return Gcd(y, x%y)
}
