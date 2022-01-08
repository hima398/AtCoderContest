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

	const nm = int(1e4)
	n, k := nextInt(), nextInt()
	d := nextIntSlice(k)

	i := n
	ok := false
	for !ok {
		ns := strconv.Itoa(i)
		ok = true
		for j := range ns {
			for _, di := range d {
				ok = ok && ns[j] != byte(di)+'0'
			}
		}
		if ok {
			fmt.Println(i)
			return
		}
		i++
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
