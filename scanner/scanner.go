package scanner

import (
	"bufio"
	"os"
	"strconv"
)

type Scanner struct {
	sc *bufio.Scanner
}

func New() *Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (sc *Scanner) NextInt() int {
	sc.sc.Scan()
	i, err := strconv.Atoi(sc.sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (sc *Scanner) NextFloat64() float64 {
	sc.sc.Scan()
	f, err := strconv.ParseFloat(sc.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func (sc *Scanner) NextString() string {
	sc.sc.Scan()
	return sc.sc.Text()
}
