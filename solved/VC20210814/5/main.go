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

	k := nextInt()

	a := make([]int, k+1)
	a[0] = 7 % k
	for i := 0; i < k; i++ {
		if a[i] == 0 {
			fmt.Println(i + 1)
			return
		}
		a[i+1] = (a[i]*10 + 7) % k
	}
	if a[k] == 0 {
		fmt.Println(k)
	} else {
		fmt.Println(-1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
