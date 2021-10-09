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

	c := make([][]string, 4)
	for i := 0; i < 4; i++ {
		c[i] = make([]string, 4)
		for j := 0; j < 4; j++ {
			c[i][j] = nextString()
		}
	}
	for i := 3; i >= 0; i-- {
		//fmt.Println(c[i])
		for j := 3; j >= 0; j-- {
			fmt.Printf("%s ", string(c[i][j]))
		}
		fmt.Println("")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
