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

	n := nextInt()
	p := nextIntSlice(n)
	m := make([]bool, 200001)
	var ans []int
	v := 0
	for _, pi := range p {
		m[pi] = true
		for m[v] {
			v++
		}
		ans = append(ans, v)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, a := range ans {
		fmt.Fprintln(out, a)
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
