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

	const p = int(1e9) + 7
	n := nextInt()
	a := nextIntSlice(n)
	if a[0] != 0 {
		fmt.Println(0)
		return
	}
	//各頂点の深さのマップ
	mxd := 0
	//md := make(map[int]int)
	md := make([]int, n)
	for _, ai := range a {
		md[ai]++
		mxd = Max(mxd, ai)
	}
	//fmt.Println(md)
	//根が2つあり得る
	if md[0] != 1 {
		fmt.Println(0)
		return
	}
	//深さiの中に頂点がないケース
	//for i := 0; i < mxd; i++ {
	//	if md[i] == 0 {
	//		fmt.Println(0)
	//		return
	//	}
	//}
	//comb := NewCombination(n, p)
	ans := 1
	for i := 1; i <= mxd; i++ {
		// 深さiの頂点から深さi-1の頂点(md[i-1])への辺の結び方のうち全く結ばない1通りを引く
		// それを深さiの頂点数分かける
		b := Pow(2, md[i-1], p) - 1

		ans *= Pow(b, md[i], p)
		ans %= p
		// 深さiの中で辺を繋ぐ数
		// 組み合わせを決めて、その各々に対して辺を結ぶか結ばないの2通り
		nc2 := md[i] * (md[i] - 1) / 2
		ans *= Pow(2, nc2, p)
		ans %= p
	}
	//fmt.Println(f)
	if ans < 0 {
		ans += p
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = (ret * x) % p
		}
		y >>= 1
		x = (x * x) % p
	}
	return ret % p
}
