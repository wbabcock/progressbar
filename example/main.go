package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/wbabcock/progressbar"
	"github.com/wbabcock/progressbar/colors"
	"github.com/wbabcock/progressbar/indicator"
)

func main() {
	p := progressbar.NewProgressBar().
		WithLength(80).
		WithIndicator(indicator.Block).
		WithColor(colors.Green)

	for i := 0; i <= 100; i++ {
		p.Update(uint(i))
		time.Sleep(time.Duration(rand.IntN(200)) * time.Millisecond)
	}
	fmt.Printf("\n")
}
