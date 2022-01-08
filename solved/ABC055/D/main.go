package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Eval(pre, i int, s byte) int {
	if i == 0 {
		if s == 'o' {
			return pre ^ 1
		} else {
			return pre
		}
	} else {
		if s == 'o' {
			return pre
		} else {
			return pre ^ 1
		}
	}
}

func ConvertAns(c []int) string {
	m := []string{"W", "S"}
	ans := ""
	for i := 0; i < len(c)-2; i++ {
		ans += m[c[i]]
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	s += string(s[0])
	s += string(s[1])

	c := make([]int, n+2)
	c[0], c[1] = 0, 0
	for i := 1; i <= n; i++ {
		c[i+1] = Eval(c[i-1], c[i], s[i])
	}
	if c[0] == c[n] && c[1] == c[n+1] {
		fmt.Println(ConvertAns(c))
		return
	}
	c = make([]int, n+2)
	c[0], c[1] = 0, 1
	for i := 1; i <= n; i++ {
		c[i+1] = Eval(c[i-1], c[i], s[i])
	}
	if c[0] == c[n] && c[1] == c[n+1] {
		fmt.Println(ConvertAns(c))
		return
	}
	c = make([]int, n+2)
	c[0], c[1] = 1, 0
	for i := 1; i <= n; i++ {
		c[i+1] = Eval(c[i-1], c[i], s[i])
	}
	if c[0] == c[n] && c[1] == c[n+1] {
		fmt.Println(ConvertAns(c))
		return
	}
	c = make([]int, n+2)
	c[0], c[1] = 1, 1
	for i := 1; i <= n; i++ {
		c[i+1] = Eval(c[i-1], c[i], s[i])
	}
	if c[0] == c[n] && c[1] == c[n+1] {
		fmt.Println(ConvertAns(c))
		return
	}
	fmt.Println(-1)
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
