package main

import (
	"fmt"

	"github.com/tv42/seed"
)

func main() {
	d6 := seed.Rand().Int31n(6)
	fmt.Println(d6)
}
