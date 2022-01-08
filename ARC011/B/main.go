package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func Convert(w string) (ret string) {
	t := map[byte]string{'b': "1", 'c': "1", 'd': "2", 'w': "2", 't': "3", 'j': "3", 'f': "4", 'q': "4", 'l': "5", 'v': "5", 's': "6", 'x': "6", 'p': "7", 'm': "7", 'h': "8", 'k': "8", 'n': "9", 'g': "9", 'z': "0", 'r': "0"}

	lw := strings.ToLower(w)
	for i := range lw {
		if _, ok := t[lw[i]]; ok {
			ret += t[lw[i]]
		}
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	w := make([]string, n)
	for i := 0; i < n; i++ {
		w[i] = nextString()
	}
	//fmt.Println(w)
	var ans []string
	for _, wi := range w {
		a := Convert(wi)
		if len(a) > 0 {
			ans = append(ans, a)
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if len(ans) > 0 {
		fmt.Fprintf(out, "%s", ans[0])
	}
	for i := 1; i < len(ans); i++ {
		fmt.Fprintf(out, " %s", ans[i])
	}
	fmt.Fprintln(out)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
