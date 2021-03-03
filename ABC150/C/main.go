package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func DeleteElement(n []string, idx int) []string {
	var ret []string
	ret = append(ret, n[:idx]...)
	ret = append(ret, n[idx+1:]...)
	return ret
}

func Permute(pat, rem []string) []string {
	var ret []string

	var dfs func(pat, rem []string)
	dfs = func(pat, rem []string) {
		if len(rem) == 0 {
			ret = append(ret, strings.Join(pat, ""))
			return
		}
		for i := 0; i < len(rem); i++ {
			dfs(append(pat, rem[i]), DeleteElement(rem, i))
		}
	}
	dfs([]string{}, rem)
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ln := make([]string, n)
	for i := 0; i < n; i++ {
		ln[i] = strconv.Itoa(i + 1)
	}

	s := Permute([]string{}, ln)

	p := make([]string, n)
	q := make([]string, n)

	for i := 0; i < n; i++ {
		p[i] = strconv.Itoa(nextInt())
	}
	for i := 0; i < n; i++ {
		q[i] = strconv.Itoa(nextInt())
	}
	sp := strings.Join(p, "")
	sq := strings.Join(q, "")
	a, b := 0, 0
	for i, v := range s {
		if v == sp {
			a = i + 1
		}
		if v == sq {
			b = i + 1
		}
		if a > 0 && b > 0 {
			break
		}
	}
	fmt.Println(Abs(a - b))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
