package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k, p := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	hn := Ceil(n, 2)
	a1 := a[:hn]
	a2 := a[hn:]
	p1 := (1 << len(a1)) - 1
	p2 := (1 << len(a2)) - 1
	//fmt.Printf("%b %b\n", p1, p2)
	s1 := make([][]int, len(a1)+1)
	s2 := make([][]int, len(a2)+1)
	for f := 0; f <= p1; f++ {
		ss := 0
		for i := 0; i < len(a1); i++ {
			if (f>>i)&1 > 0 {
				ss += a1[i]
			}
		}
		s1[bits.OnesCount(uint(f))] = append(s1[bits.OnesCount(uint(f))], ss)
	}
	for f := 0; f <= p2; f++ {
		ss := 0
		for i := 0; i < len(a2); i++ {
			if (f>>i)&1 > 0 {
				ss += a2[i]
			}
		}
		s2[bits.OnesCount(uint(f))] = append(s2[bits.OnesCount(uint(f))], ss)
	}

	//fmt.Println(a1)
	//fmt.Println(a2)
	// s1とs2でインデックスのビットがk、合計がp以下の数
	for _, s := range s1 {
		sort.Ints(s)
	}
	for _, s := range s2 {
		sort.Ints(s)
	}
	//fmt.Println(s1)
	//fmt.Println(s2)
	ans := 0
	for i := 0; i <= Min(k, len(s1)-1); i++ {
		if k-i < len(s2) {
			for _, v := range s1[i] {
				// s1からi個選んだ合計と、s2からk-i個選んだ合計との和がpを超えるs2[k-i]のインデックスを探す
				// インデックスの値が、上記のp以下になる組み合わせの個数
				idx := sort.Search(len(s2[k-i]), func(ii int) bool {
					return v+s2[k-i][ii] > p
				})
				//fmt.Println(i, v, idx)
				ans += idx
			}
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
