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
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			fmt.Println("Bad")
			return
		}
	}
	fmt.Println("Good")

}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
