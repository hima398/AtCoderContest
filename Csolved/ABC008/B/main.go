package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Candidate struct {
	n    int
	name string
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		s := nextString()
		m[s]++
	}
	var c []Candidate
	for k, v := range m {
		c = append(c, Candidate{v, k})
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i].n > c[j].n
	})
	fmt.Println(c[0].name)
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
