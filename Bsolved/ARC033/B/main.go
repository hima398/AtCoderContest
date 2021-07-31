package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	na, nb := nextInt(), nextInt()
	a := nextIntSlice(na)
	b := nextIntSlice(nb)

	s1, s2 := make(map[int]int), make(map[int]int)
	for _, v := range a {
		s1[v]++
		s2[v]++
	}
	for _, v := range b {
		s1[v]++
		s2[v]++
	}
	for k, v := range s1 {
		// a、bに同じ数字が出てこないのでaにもbにも現れる物は2になっているはず
		if v < 2 {
			delete(s1, k)
		}
	}
	ans := float64(len(s1)) / float64(len(s2))
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
