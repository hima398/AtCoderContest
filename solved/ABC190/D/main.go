package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Divisor(n int) map[int]int {
	ret := make(map[int]int)
	ret[1] = n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ret[i] = n / i
			ret[n/i] = i
		}
	}
	ret[n] = 1
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	if n <= 2 {
		fmt.Println(2)
		return
	}
	d := Divisor(2 * n)
	//fmt.Println(d)
	ans := 0
	for k, v := range d {
		//fmt.Println(k, v, v-k+1)
		if Abs(v-k)%2 == 1 {
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
