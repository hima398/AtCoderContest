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
	fmt.Printf("%b %b\n", p1, p2)
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
	fmt.Println(s1)
	fmt.Println(s2)
	ans := 0
	for i := 0; i <= k; i++ {
		idx := sort.Search(len(s1[i]), func(j int) bool {
			return s1[i][j] > p
		})
		//fmt.Println(i, idx)
		idx2 := sort.Search(len(s2[i]), func(j int) bool {
			return s1[i][idx-1]+s2[k-i][j] > p
		})
		fmt.Println(i, idx, idx2)
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

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
