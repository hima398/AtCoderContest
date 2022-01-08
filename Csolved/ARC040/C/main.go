package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println()
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	r := n
	ans := 0
	for i := 0; i < n; i++ {
		for j := r - 1; j >= 0; j-- {
			if s[i][j] == '.' {
				ans++
				r = j
				if i < n-1 {
					s[i+1] = strings.Join(append(strings.Split(s[i+1], "")[:r], strings.Split(strings.Repeat("o", n-r), "")...), "")
				}
				//PrintField(s)
				if r == 0 {
					r = n
				}
				break
			}
			if j == 0 {
				r = n
			}
		}
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
