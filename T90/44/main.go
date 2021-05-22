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

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	offset := 0
	for i := 0; i < q; i++ {
		t, x, y := nextInt(), nextInt(), nextInt()
		x--
		y--
		switch t {
		case 1:
			i := (x + offset) % n
			j := (y + offset) % n
			a[i], a[j] = a[j], a[i]
			//fmt.Fprintln(out, a)
		case 2:
			offset--
			if offset < 0 {
				offset += n
			}
			offset %= n
		case 3:
			i := (x + offset) % n
			//fmt.Fprintln(out, a)
			fmt.Fprintln(out, a[i])
		default:
		}
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
