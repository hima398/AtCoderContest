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

	n := nextInt()
	a := nextIntSlice(n)
	s := 0
	for i := 0; i < n; i++ {
		s += a[i]
		a = append(a, a[i])
	}
	//整数の和がsの10分の1が比較できない
	if s%10 != 0 {
		fmt.Println("No")
		return
	}
	s /= 10
	ps := a[0]
	l, r := 0, 1
	for {
		if s == ps {
			fmt.Println("Yes")
			return
		}
		if r < 2*n && (s > ps || l == r) { // rを進める場合
			ps += a[r]
			r++
		} else if l < r && s < ps { // lを進める場合
			ps -= a[l]
			l++
		} else {
			break
		}

	}
	fmt.Println("No")
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
