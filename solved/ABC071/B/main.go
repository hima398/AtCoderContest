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

	var t [26]bool
	s := nextString()
	for i := range s {
		t[s[i]-'a'] = true
	}
	//fmt.Println(t)
	for i, IsFound := range t {
		if !IsFound {
			fmt.Println(string(i + 'a'))
			return
		}
	}
	fmt.Println("None")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
