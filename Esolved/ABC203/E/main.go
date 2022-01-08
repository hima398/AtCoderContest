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

type Pawn struct {
	i, j int
}

func Find(s []int, x int) bool {
	/*
		for _, v := range s {
			if v == x {
				return true
			}
		}
		return false
	*/
	idx := sort.Search(len(s), func(i int) bool {
		return s[i] >= x
	})
	return idx < len(s) && s[idx] == x
}

func SolveDP() {
	n, m := nextInt(), nextInt()
	bp := make(map[int]map[int]bool)
	bpi := make([]int, m)
	for i := 0; i < m; i++ {
		ii, jj := nextInt(), nextInt()
		bpi[i] = ii
		if bp[ii] == nil {
			bp[ii] = make(map[int]bool)
		}
		bp[ii][jj] = true
	}
	sort.Ints(bpi)
	/*
		for i := range bp {
			sort.Ints(bp[i])
		}
	*/

	dp := make(map[int]map[int]bool)
	dp[0] = make(map[int]bool)
	dp[0][n] = true
	idx := 0
	for _, ni := range bpi {
		if dp[ni] == nil {
			dp[ni] = make(map[int]bool)
		}
		for k, ok := range dp[idx] {
			if ok {
				njl, nj, njr := k-1, k, k+1
				if njl >= 0 {
					dp[ni][njl] = true
				}
				dp[ni][nj] = true
				if njr <= 2*n {
					dp[ni][njr] = true
				}
			}
		}
		for j := range dp[ni] {
			if _, ok := bp[ni][j]; !ok {
				dp[ni][j] = false
			}
		}
		idx = ni
	}
	fmt.Println(dp[2*n])
	ans := 0
	for _, ok := range dp[2*n] {
		if ok {
			ans++
		}
	}

	fmt.Println(ans)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var bp []Pawn
	for i := 0; i < m; i++ {
		p := Pawn{nextInt(), nextInt()}
		bp = append(bp, p)
	}
	sort.Slice(bp, func(i, j int) bool {
		if bp[i].i < bp[j].i {
			return true
		}
		if bp[i].i == bp[j].i {
			return bp[i].j < bp[j].j
		}
		return false
	})
	s := make(map[int]bool)
	s[n] = true
	idx := 0
	var add, del []int
	for idx < m {
		add, del = []int{}, []int{}
		next := idx
		for next < m {
			if bp[idx].i == bp[next].i {
				next++
			} else {
				break
			}
		}
		for i := idx; i < next; i++ {
			jj := bp[i].j
			if (s[jj-1] || s[jj+1]) && !s[jj] {
				add = append(add, jj)
			}
			if !s[jj-1] && !s[jj+1] && s[jj] {
				del = append(del, jj)
			}
		}
		for _, v := range add {
			s[v] = true
		}
		for _, v := range del {
			s[v] = false
		}
		idx = next
	}
	//fmt.Println(s)
	ans := 0
	for _, ok := range s {
		if ok {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
