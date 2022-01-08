package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34
	//c o n t e s t c o  n  t  e  s  t  c  o  n  t  e  s  t  c  o  n  t  e  s  t  c  o  n  t  e  st
	//s  e  n  t  e  n  c  e
	//6 12 17 18 19 24 29 33
	s := nextString()
	t := nextString()
	dic := make([][]int, 26)
	for i := range s {
		idx := int(s[i] - 'a')
		dic[idx] = append(dic[idx], i)
		dic[idx] = append(dic[idx], i+len(s))
	}
	for i := 0; i < 26; i++ {
		sort.Ints(dic[i])
	}
	idx := 0
	ans := big.NewInt(0)
	for i := range t {
		c := t[i] - 'a'
		if len(dic[int(c)]) == 0 {
			fmt.Println(-1)
			return
		}
		idx = sort.Search(len(dic[int(c)]), func(k int) bool {
			return idx <= dic[c][k]
		})
		idx = dic[c][idx] + 1
		//fmt.Println(string(t[i]), dic[int(c)], i, idx)
		if idx >= len(s) {
			idx -= len(s)
			ans = ans.Add(ans, big.NewInt(int64(len(s))))
		}
	}
	ans = ans.Add(ans, big.NewInt(int64(idx)))
	fmt.Println(ans.String())
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
