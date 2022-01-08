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
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	var ans int
	var ra, ab, fb int
	for _, ss := range s {
		ans += strings.Count(ss, "AB")

		if ss[0] == 'B' && ss[len(ss)-1] == 'A' {
			ab++
		} else if ss[0] == 'B' {
			fb++
		} else if ss[len(ss)-1] == 'A' {
			ra++
		}
	}
	if ab == 0 {
		ans += Min(ra, fb)
	} else {
		if ra+fb > 0 {
			ans += ab + Min(ra, fb)
		} else {
			ans += ab - 1
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
