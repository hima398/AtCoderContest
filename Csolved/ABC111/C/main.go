package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Node struct {
	k, v int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	v := nextIntSlice(n)

	//全て同じ
	IsSame := true
	for i := 0; i < n-1; i++ {
		IsSame = IsSame && (v[i+1] == v[i])
	}
	if IsSame {
		fmt.Println(n / 2)
		return
	}

	// 偶奇の中でそれぞれ一番多い数字を残して書き換える
	me, mo := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			me[v[i]]++
		} else {
			mo[v[i]]++
		}
	}
	var se, so []Node
	for k, v := range me {
		se = append(se, Node{k, v})
	}
	for k, v := range mo {
		so = append(so, Node{k, v})
	}
	sort.Slice(se, func(i, j int) bool { return se[i].v > se[j].v })
	sort.Slice(so, func(i, j int) bool { return so[i].v > so[j].v })
	var ans int

	if se[0].k != so[0].k {
		ans = n - se[0].v - so[0].v
	} else {
		ans = Min(n-se[0].v-so[1].v, n-se[1].v-so[0].v)
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
