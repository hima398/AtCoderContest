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

	n, m := nextInt(), nextInt()
	s := nextIntSlice(n)
	t := nextIntSlice(m)

	// tが作成可能かチェックする
	ms, mt := make(map[int]int), make(map[int]int)
	for _, v := range s {
		ms[v]++
	}
	for _, v := range t {
		mt[v]++
	}
	// sが0, 1の1種類しかないケースを先に処理しておく
	if len(ms) == 1 {
		// tに必要な要素がsにない場合はtを作成不可能
		for k := range mt {
			if _, ok := ms[k]; !ok {
				fmt.Println(-1)
				return
			}
		}
		// そうでなければsとtは同じ数字が並んでいるだけなので
		// シフトせずm回操作するのが答え
		fmt.Println(m)
		return
	}

	//作業用のスライス
	ss := append(s, s...)
	lz, lo, rz, ro := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	d := 0
	// 左に一番近い0
	// e1であれば
	// 0 0 1
	for i := 0; i < 2*n-1; i++ {
		lz[i%n] = d
		if ss[i+1] == 0 {
			d = 0
		} else {
			d++
		}
	}
	d = 0
	// 左に一番近い1
	for i := 0; i < 2*n-1; i++ {
		lo[i%n] = d
		if ss[i+1] == 1 {
			d = 0
		} else {
			d++
		}
	}
	// 右に一番近い0
	for i := 2*n - 1; i > 0; i-- {
		rz[i%n] = d
		if ss[i-1] == 0 {
			d = 0
		} else {
			d++
		}
	}
	d = 0
	// 右に一番近い1
	for i := 2*n - 1; i > 0; i-- {
		ro[i%n] = d
		if ss[i-1] == 1 {
			d = 0
		} else {
			d++
		}
	}

	//fmt.Println(lz, lo)
	//fmt.Println(rz, ro)
	// 右に一番近い1
	// 2 1 0
	// 左に一番近い1
	// 1 2 0
	// 右に一番近い0
	// 左に一番近い0
	// 0 0 1
	//入力1の例
	//l := [][]int{{0, 0, 1}, {1, 2, 0}}
	//r := [][]int{{0, 0, 1}, {2, 1, 0}}
	l := [][]int{lz, lo}
	r := [][]int{rz, ro}
	idx := 0
	//最低len(t) = m回は操作する必要がある
	ans := m
	for i := 0; i < m; i++ {
		ld, rd := l[t[i]][idx], r[t[i]][idx]
		if ld <= rd {
			ans += ld
			idx = (idx - ld + n) % n
		} else {
			ans += rd
			idx = (idx + rd) % n
		}
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
