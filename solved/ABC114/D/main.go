package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Divisor(n int) []int {
	ret := []int{1, n}
	if n%2 == 0 {
		ret = append(ret, 2)
		ret = append(ret, n/2)
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i != n/i {
				ret = append(ret, n/i)
			}
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	e := make([]int, n+1)
	for i := 2; i <= n; i++ {
		cur := i
		for j := 2; j <= i; j++ {
			if cur <= 1 {
				break
			}
			for cur%j == 0 {
				e[j]++
				cur /= j
			}
		}
	}
	//fmt.Println(e)
	Count := func(n int) int {
		ret := 0
		for _, v := range e {
			if v >= n-1 {
				ret++
			}
		}
		return ret
	}
	ans := Count(75) + Count(25)*(Count(3)-1) + (Count(15) * (Count(5) - 1)) + (Count(5) * (Count(5) - 1) * (Count(3) - 2) / 2)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
