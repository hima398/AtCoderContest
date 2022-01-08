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

	e := nextIntSlice(6)
	b := nextInt()
	l := nextIntSlice(6)
	me := make(map[int]bool)
	for _, v := range e {
		me[v] = true
	}
	n := 0 // 一致した数
	IsBonus := false
	for _, v := range l {
		if me[v] {
			n++
		}
		if v == b {
			IsBonus = true
		}
	}
	if n == 6 {
		fmt.Println(1)
	} else if n == 5 {
		if IsBonus {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	} else if n == 4 {
		fmt.Println(4)
	} else if n == 3 {
		fmt.Println(5)
	} else {
		fmt.Println(0)
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
