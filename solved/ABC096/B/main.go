package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := make([]int, 3)
	s[0], s[1], s[2] = nextInt(), nextInt(), nextInt()
	k := nextInt()
	sort.Ints(s)
	ans := s[0] + s[1] + s[2]
	for i := 0; i < k; i++ {
		ans += s[2]
		s[2] *= 2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
