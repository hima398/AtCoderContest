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
	if s == "Sunny" {
		fmt.Println("Cloudy")
	} else if s == "Cloudy" {
		fmt.Println("Rainy")
	} else {
		// Rainy
		fmt.Println("Sunny")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
