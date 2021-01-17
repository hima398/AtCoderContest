package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrimeFactor(n int) map[int]int {
	ans := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			ans[i]++
			n /= i
		}
	}
	if n != 1 {
		ans[n]++
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := PrimeFactor(n)

	//fmt.Println(m)
	ans := 0
	for _, v := range a {
		t := 1
		i := 1
		for v-t >= 0 {
			ans++
			i++
			t += i
		}
	}

	fmt.Println(ans)
}
