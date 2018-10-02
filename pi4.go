package  main

import (
	"fmt"
	"time"
)

const N = 100000000
const TH = 4

var ch chan float64

func calcPi() float64 {
	step := 1.0 / N

	for i := 0; i < TH; i++ {
		go func(idx int) {
			sum := 0.0
			for k := idx; k < N; k += TH {
				x := (float64(k) + 0.5)*step
				sum += 4.0 / (1.0 + x*x)
			}
			ch <- sum*step
		}(i)
	}
	pi := 0.0
	for i := 0; i < TH; i++ {
		pi += <-ch
	}

	return pi
}

func main() {
	ch = make(chan float64)
	inicio := time.Now()
	fmt.Println(calcPi())
	fmt.Printf("Tiempo: %s\n", time.Since(inicio))
}