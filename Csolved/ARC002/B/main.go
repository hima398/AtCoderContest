package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	ss := nextString()
	t, _ := time.Parse("2006/01/02", ss)

	for {
		if t.Year()%(int(t.Month())*t.Day()) == 0 {
			fmt.Println(t.Format("2006/01/02"))
			return
		}
		t = t.AddDate(0, 0, 1)
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
