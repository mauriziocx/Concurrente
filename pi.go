package  main

import (
	"fmt"
	"time"
)

const N = 100000000

func calcPi() float64 {
	var x, sum float64
	sum = 0.0

	step := 1.0 / N

	for i := 0; i < N; i++ {
		x = (float64(i) + 0.5)*step
		sum += 4 / (1.0 + x*x)
	}
	return step * sum
}

func main() {
	inicio := time.Now()
	fmt.Println(calcPi())
	fmt.Printf("Tiempo: %s\n", time.Since(inicio))
}