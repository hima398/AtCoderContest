package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := strings.Split(nextString(), "")
	var ans []string
	for i := range s {
		switch s[i] {
		case "0":
			ans = append(ans, "0")
		case "1":
			ans = append(ans, "1")
		case "B":
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		}
		//fmt.Println(ans)
	}
	fmt.Println(strings.Join(ans, ""))
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
