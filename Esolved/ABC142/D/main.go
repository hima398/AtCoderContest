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

func PrimeFactorization(n int) map[int]int {

	ret := make(map[int]int)
	for i := 2; i <= rootMax; i++ {
		if !soe[i] {
			for n%i == 0 {
				ret[i]++
				n /= i
			}
		}
		if n == 1 {
			break
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
	pfa := PrimeFactorization(a)
	pfb := PrimeFactorization(b)
	//fmt.Println(pfa)
	//fmt.Println(pfb)
	//fmt.Println(soe)
	ans := 1
	for k := range pfa {
		if pfb[k] > 0 {
			ans++
		}
	}

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
