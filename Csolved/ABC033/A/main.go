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

	n := nextString()
	IsSame := true
	for i := 1; i < 4; i++ {
		IsSame = IsSame && (n[0] == n[i])
	}
	if IsSame {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
