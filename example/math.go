package main

import (
	"fmt"

	"github.com/syumai/proconutil/math"
)

func main() {
	fmt.Println(math.Max(3, 1, 2))
	fmt.Println(math.Min(3, 1, 2))

	fmt.Println(math.GCD(18, 30))
}
