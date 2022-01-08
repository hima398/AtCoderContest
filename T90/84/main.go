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

	n := nextInt()
	s := nextString()
	cnts := make([]int, n)
	now := s[n-1]
	for i := n - 2; i >= 0; i-- {
		if s[i] == now {
			if cnts[i+1] > 0 {
				cnts[i] = cnts[i+1] + 1
			} else {
				cnts[i] = 0
			}
		} else {
			cnts[i] = 1
			now = s[i]
		}
	}
	//fmt.Println(cnts)
	ans := 0
	for i, cnt := range cnts {
		if cnt == 0 {
			continue
		}
		ans += n - (cnt + i)
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
