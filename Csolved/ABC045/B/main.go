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

	const NP = 3 // num players
	players := make([]string, NP)
	for i := 0; i < NP; i++ {
		players[i] = nextString()
	}
	next := 0
	for {
		if len(players[next]) == 0 {
			break
		}
		c := players[next][0]
		players[next] = players[next][1:]
		next = int(c - 'a')
	}
	ans := string(next + 'A')
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
