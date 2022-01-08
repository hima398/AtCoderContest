package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Word struct {
	s, r string
}

func Reverse(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		r += string(s[len(s)-1-i])
	}
	return r
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//ss := make([]string, n)
	ss := make([]Word, n)
	for i := 0; i < n; i++ {
		ss[i].s = nextString()
		ss[i].r = Reverse(ss[i].s)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].r < ss[j].r
	})
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ss[i].s)
	}
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
