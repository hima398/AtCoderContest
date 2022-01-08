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
	s, t := nextString(), nextString()
	ss := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		//fmt.Println(ss)
		if strings.Join(ss, "") == t {
			fmt.Println("Yes")
			return
		}
		ss = append([]string{ss[len(ss)-1]}, ss[:len(ss)-1]...)
	}
	fmt.Println("No")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
