package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	p := nextIntSlice(n - 1)
	e := make(map[int][]int)
	for i, pi := range p {
		pi--
		e[i+1] = append(e[i+1], pi)
		e[pi] = append(e[pi], i+1)
	}
	//fmt.Println(e)
	s := make([]int, n)
	t := make([]int, n)
	md := make(map[int][]int)
	var Dfs func(x, p, d, r int) int
	Dfs = func(x, p, d, r int) int {
		s[x] = r
		o := r
		md[d] = append(md[d], s[x])
		for _, v := range e[x] {
			if v != p {
				o = Dfs(v, x, d+1, o+1)
			}
		}
		o++
		t[x] = o
		return o
	}
	Dfs(0, -1, 0, 0)
	//fmt.Println(s)
	//fmt.Println(t)
	//fmt.Println(md)
	for k := range md {
		sort.Ints(md[k])
	}
	q := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		u, d := nextInt(), nextInt()
		u--
		//fmt.Println(u, d, s[u], t[u], md[d])
		idx1 := sort.Search(len(md[d]), func(i int) bool {
			return md[d][i] > t[u]
		})
		idx2 := sort.Search(len(md[d]), func(i int) bool {
			return s[u] <= md[d][i]
		})
		//fmt.Println(idx1, idx2)
		ans := idx1 - idx2
		fmt.Fprintln(out, ans)
	}
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
