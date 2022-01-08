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
	a := nextIntSlice(n)
	//fmt.Println(a)
	set := make(map[int]bool)
	for _, ai := range a {
		set[ai] = true
	}
	var m []int
	for k := range set {
		m = append(m, k)
	}
	sort.Ints(m)
	t := make(map[int]int)
	for i, v := range m {
		t[v] = i
	}
	var b []int
	for _, ai := range a {
		b = append(b, t[ai])
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range b {
		fmt.Fprintln(out, v)
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
