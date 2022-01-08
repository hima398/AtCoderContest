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
	n := len(s)
	var ans string
	m := map[byte]string{'O': "0", 'D': "0", 'I': "1", 'Z': "2", 'S': "5", 'B': "8"}
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; !ok {
			ans += string(s[i])
		} else {
			ans += m[s[i]]
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
