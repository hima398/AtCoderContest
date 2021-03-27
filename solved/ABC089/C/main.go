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
	//m := make(map[byte]int)
	e := map[byte]int{'M': 0, 'A': 1, 'R': 2, 'C': 3, 'H': 4}
	m := [5]int{}
	for i := 0; i < n; i++ {
		si := nextString()
		c := si[0]
		_, ok := e[c]
		if ok {
			m[e[c]]++
		}
	}
	ans := 0
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				ans += m[i] * m[j] * m[k]
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
