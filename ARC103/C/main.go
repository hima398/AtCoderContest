package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Item struct {
	K int
	V int
}

type List []Item

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return l[i].V > l[j].V
}

func (l List) Print() {
	for i := 0; i < l.Len(); i++ {
		fmt.Printf("K:%d, V:%d\n", l[i].K, l[i].V)
	}

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	mOdd := make(map[int]int)
	mEven := make(map[int]int)
	for i := 1; i <= n; i++ {
		v := nextInt()
		if i%2 == 1 {
			mOdd[v]++
		} else {
			mEven[v]++
		}
	}
	//fmt.Println(mOdd)
	//fmt.Println(mEven)
	var lOdd List
	var lEven List
	for k, v := range mOdd {
		i := Item{k, v}
		lOdd = append(lOdd, i)
	}
	for k, v := range mEven {
		i := Item{k, v}
		lEven = append(lEven, i)
	}
	sort.Sort(lOdd)
	sort.Sort(lEven)

	//lOdd.Print()
	//lEven.Print()
	o1 := lOdd[0]
	e1 := lEven[0]
	var ans int
	if o1.K == e1.K {
		if lOdd.Len() == 1 || lEven.Len() == 1 {
			ans = Min(lOdd[0].V, lEven[0].V)
		} else {
			o2 := lOdd[1]
			e2 := lEven[1]
			ans = Min(n-o1.V-e2.V, n-o2.V-e1.V)
		}
	} else {
		ans = n - o1.V - e1.V
	}
	fmt.Println(ans)
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
