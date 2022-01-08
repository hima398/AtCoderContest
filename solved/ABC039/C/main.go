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
	t := "WBWBWWBWBWBW"
	t += t + t
	m := []string{
		0: "Do", 1: "Do", 2: "Re", 3: "Re",
		4: "Mi", 5: "Fa", 6: "Fa", 7: "So",
		8: "So", 9: "La", 10: "La", 11: "Si"}
	for i := 0; i < len(m); i++ {
		ok := true
		for j := i; j < i+len(s); j++ {
			if t[j] != s[j-i] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(m[i])
			return
		}
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
