package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	w := nextString()
	m := make(map[rune]int)
	for _, r := range w {
		m[r]++
	}
	IsBeautiful := true
	for _, v := range m {
		if v%2 == 1 {
			IsBeautiful = false
			break
		}
	}
	if IsBeautiful {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
