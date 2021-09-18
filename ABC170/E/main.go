package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n, q int, a, b, c, d []int) []int {
	const INF = int(1e9) + 1
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		//転園
		b[c[i]] = d[i]
		m := make(map[int]int)
		for j := 0; j < n; j++ {
			m[b[j]] = Max(m[b[j]], a[j])
		}
		ans[i] = INF
		for _, v := range m {
			ans[i] = Min(ans[i], v)
		}
	}
	return ans
}

//平行二分木を使ってシミュレート
func Solve(n, q int, a, b, c, d []int) []int {
	ans := make([]int, q)

	IncrementTree := func(k int, t *redblacktree.Tree) {
		v, found := t.Get(k)
		if found {
			t.Put(k, v.(int)+1)
		} else {
			t.Put(k, 1)
		}
	}

	DecrementTree := func(k int, t *redblacktree.Tree) {
		v, found := t.Get(k)
		if found {
			if v.(int) > 1 {
				t.Put(k, v.(int)-1)
			} else if v == 1 {
				t.Remove(k)
			}
		}
	}

	km := make(map[int]*redblacktree.Tree)
	for i := 0; i < n; i++ {
		if _, ok := km[b[i]]; !ok {
			km[b[i]] = redblacktree.NewWithIntComparator()
		}
		IncrementTree(a[i], km[b[i]])
	}
	ms := redblacktree.NewWithIntComparator()
	for k := range km {
		max := km[k].Right().Key.(int)
		IncrementTree(max, ms)
	}

	for i := 0; i < q; i++ {
		key := b[c[i]] //c[i]のが所属している幼稚園

		bm1 := km[key].Right().Key.(int) //c[i]が所属している幼稚園から転出前の最大値
		DecrementTree(a[c[i]], km[key])
		am1 := -1
		if km[key].Size() > 0 {
			am1 = km[key].Right().Key.(int) //c[i]が所属している幼稚園から転出後の最大値
		}

		b[c[i]] = d[i]

		if _, ok := km[d[i]]; !ok {
			km[d[i]] = redblacktree.NewWithIntComparator()
		}

		bm2 := -1
		if km[d[i]].Size() > 0 {
			bm2 = km[d[i]].Right().Key.(int) //幼稚園d[i]に転入前の最大値
		}
		IncrementTree(a[c[i]], km[d[i]])
		am2 := km[d[i]].Right().Key.(int) //幼稚園d[i]に転入後の最大値

		DecrementTree(bm1, ms)
		if am1 >= 0 {
			IncrementTree(am1, ms)
		}
		if bm2 >= 0 {
			DecrementTree(bm2, ms)
		}
		IncrementTree(am2, ms)
		ans[i] = ms.Left().Key.(int)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		b[i]-- //0-indexed
	}
	c, d := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		c[i], d[i] = nextInt(), nextInt()
		c[i]-- //0-indexed
		d[i]-- //0-indexed
	}

	//ans := SolveHonestly(n, q, a, b, c, d)
	ans := Solve(n, q, a, b, c, d)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
