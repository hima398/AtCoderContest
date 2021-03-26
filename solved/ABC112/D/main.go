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
	ret := []int{1}
	if n%2 == 0 {
		ret = append(ret, 2)
		ret = append(ret, n/2)
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			ret = append(ret, n/i)
		}
	}
	ret = append(ret, n)
	sort.Ints(ret)
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	d := Divisor(m)
	//fmt.Println(n, m, d)
	ans := 0
	for _, v := range d {
		if n*v <= m {
			ans = v
		} else {
			break
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
