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

	a, b, c := true, false, true

	if a {
		fmt.Printf("At")
	} else {
		fmt.Printf("Yo")
	}

	if !a && b {
		fmt.Printf("Bo")
	} else if !b || c {
		fmt.Printf("Co")
	}

	if a && b && c {
		fmt.Printf("foo!")
	} else if true && false {
		fmt.Printf("yeah!")
	} else if !a || c {
		fmt.Printf("der")
	}
	fmt.Println()
}
