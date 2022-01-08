package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	var min [26]int
	for i := 0; i < 26; i++ {
		min[i] = 51
	}
	for i := 0; i < n; i++ {
		var m [26]int
		for j := 0; j < len(s[i]); j++ {
			m[s[i][j]-'a']++
		}
		for k, v := range m {
			min[k] = Min(min[k], v)
		}
	}
	ans := ""
	for i, v := range min {
		ans += strings.Repeat(string('a'+byte(i)), v)
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
