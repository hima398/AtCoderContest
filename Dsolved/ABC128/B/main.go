package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Restaurant struct {
	number int
	city   string
	score  int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	r := make([]Restaurant, n)
	for i := 0; i < n; i++ {
		s, p := nextString(), nextInt()
		r[i] = Restaurant{i + 1, s, p}
	}
	sort.Slice(r, func(i, j int) bool {
		if r[i].city < r[j].city {
			return true
		} else if r[i].city == r[j].city {
			return r[i].score > r[j].score
		}
		return false
	})
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range r {
		fmt.Fprintln(out, v.number)
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
