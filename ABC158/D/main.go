package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func Reverse(s string) string {
	n := len(s)
	r := make([]string, n)
	for i, _ := range s {
		r[n-1-i] = string(s[i])
	}
	return strings.Join(r, "")
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	q := nextInt()
	var lb, rb strings.Builder
	IsReversed := false
	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			IsReversed = !IsReversed
		case 2:
			f, c := nextInt(), nextString()
			switch f {
			case 1:
				if IsReversed {
					rb.WriteString(c)
				} else {
					lb.WriteString(c)
				}
			case 2:
				if IsReversed {
					lb.WriteString(c)
				} else {
					rb.WriteString(c)
				}
			}
		}
	}
	ans := Reverse(lb.String()) + s + rb.String()
	if IsReversed {
		// 反転して出力
		fmt.Println(Reverse(ans))
	} else {
		fmt.Println(ans)
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
