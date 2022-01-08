package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const p = int(1e9) + 7

func ComputeFibonacci(n int) []int {
	ret := []int{1, 1}
	n--
	if n < 1 {
		return ret
	}
	// n >= 1
	for i := 0; i < n; i++ {
		v := ret[len(ret)-1] + ret[len(ret)-2]
		v %= p
		ret = append(ret, v)
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	f := ComputeFibonacci(n)
	//fmt.Println(f[len(f)-1])
	a := nextIntSlice(n)

	nf := len(f)
	ans := f[nf-1] * a[0]
	//fmt.Println(ans)
	for i := 1; i < n; i++ {
		//fmt.Println("np: ", f[nf-i-1], f[i], f[nf-i-1]*f[i])
		//fmt.Println("nm: ", f[nf-i-2], f[i-1], f[nf-i-2]*f[i-1])
		np := f[nf-i-1] * f[i]
		nm := f[nf-i-2] * f[i-1]
		w := np - nm
		if w < 0 {
			w += p
		}
		w %= p
		ans += w * a[i]
		ans %= p
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
