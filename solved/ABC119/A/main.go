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

	s := nextString()
	var y, m, d int
	fmt.Sscanf(s, "%d/%d/%d", &y, &m, &d)
	//fmt.Println(y, m, d)
	if y <= 2019 {
		if m <= 4 {
			if d <= 30 {
				fmt.Println("Heisei")
				return
			}
		}
	}
	fmt.Println("TBD")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
