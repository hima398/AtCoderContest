package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = int(1e10) + 1
	n, m := nextInt(), nextInt()
	h := nextIntSlice(n)
	w := nextIntSlice(m)
	sort.Ints(h)
	sort.Ints(w)
	sl, sr := make([]int, (n+1)/2), make([]int, (n+1)/2)
	for i := 0; i < n/2; i++ {
		sl[i+1] = h[2*i+1] - h[2*i]
		// h[1] - h[0], h[3] - h[2] ....
	}
	for i := 1; i < len(sl); i++ {
		sl[i] += sl[i-1]
	}
	for i := 0; i < n/2; i++ {
		sr[i] = h[2*i+2] - h[2*i+1]
		// h[2] - h[1], h[4] - h[3] ....
	}
	for i := len(sr) - 2; i >= 0; i-- {
		sr[i] += sr[i+1]
	}

	//fmt.Println(sl)
	//fmt.Println(sr)
	ans := INF
	for _, v := range w {
		idx := LowerBound(h, v)
		ds := sl[idx/2] + Abs(v-h[idx/2*2]) + sr[idx/2]
		ans = Min(ans, ds)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func LowerBound(s []int, t int) int {
	idx := sort.Search(len(s), func(i int) bool {
		return s[i] >= t
	})
	return idx
}
