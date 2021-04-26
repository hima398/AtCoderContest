package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func Offset(i, n int) int {
	if i <= n {
		i += n
	} else {
		i -= n
	}
	i--
	return i
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	ss := strings.Split(s, "")
	q := nextInt()
	flag := 0
	for i := 0; i < q; i++ {
		t, a, b := nextInt(), nextInt(), nextInt()
		if t == 1 {
			if flag == 0 {
				a--
				b--
			} else {
				a = Offset(a, n)
				b = Offset(b, n)
				//fmt.Println(a, b)
			}
			ss[a], ss[b] = ss[b], ss[a]
			//fmt.Println(strings.Join(ss, ""))
		} else {
			// t == 2
			flag ^= 1
		}
	}
	if flag == 1 {
		ss = append(ss[n:], ss[:n]...)
	}
	ans := strings.Join(ss, "")
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
