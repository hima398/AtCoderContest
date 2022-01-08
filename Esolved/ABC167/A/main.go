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

	s, t := nextString(), nextString()
	for i, _ := range s {
		if s[i] != t[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
