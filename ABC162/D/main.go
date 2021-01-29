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

	_ = nextInt()
	s := nextString()
	m := make(map[string][]int)
	for i, _ := range s {
		m[string(s[i])] = append(m[string(s[i])], i+1)
	}
	ans := 0
	for _, vr := range m["R"] {
		for _, vg := range m["G"] {
			for _, vb := range m["B"] {
				v := []int{vr, vg, vb}
				sort.Ints(v)
				if v[2]-v[1] != v[1]-v[0] {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
