package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	// sc.Split(bufio.ScanWords)

	sc.Scan()
	line := sc.Text()
	regex := regexp.MustCompile(`(\s)+`)
	ans := regex.ReplaceAllString(line, ",")

	fmt.Println(ans)
}
