package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrimeFactorization(n int) (ret []int) {
	for n%2 == 0 {
		ret = append(ret, 2)
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n == 1 {
			break
		}
		for n%i == 0 {
			ret = append(ret, i)
			n /= i
		}
	}
	if n != 1 {
		ret = append(ret, n)
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ps := PrimeFactorization(n)
	//fmt.Println(len(ps), ps)
	d := 1
	ans := 0
	for len(ps) > d {
		ans++
		d *= 2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
