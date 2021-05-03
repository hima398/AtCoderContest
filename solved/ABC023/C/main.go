package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	lr, lc, k := nextInt(), nextInt(), nextInt()
	n := nextInt()
	r, c := make([]int, n), make([]int, n)
	mr, mc := make([]int, lr), make([]int, lc)
	for i := 0; i < n; i++ {
		r[i], c[i] = nextInt(), nextInt()
		r[i]--
		c[i]--
		mr[r[i]]++
		mc[c[i]]++
	}
	rcan, ccan := make([]int, n+1), make([]int, n+1)
	for i := 0; i < lr; i++ {
		rcan[mr[i]]++
	}
	for j := 0; j < lc; j++ {
		ccan[mc[j]]++
	}
	//fmt.Println(rcan)
	//fmt.Println(ccan)
	ans := 0
	for i := 0; i <= k; i++ {
		ans += rcan[i] * ccan[k-i]
	}
	//fmt.Println(ans)
	plus, minus := 0, 0
	for i := 0; i < n; i++ {
		//飴があるマス上で、行内の飴の個数+列内の飴の個数の和は答えではないので省く
		if mr[r[i]]+mc[c[i]] == k {
			minus++
		}
		//飴があるマス上で、行内の飴の個数+列内の飴の個数の和がk+1になるのは答えなので追加する
		if mr[r[i]]+mc[c[i]] == k+1 {
			plus++
		}
	}
	ans = ans - minus + plus
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
