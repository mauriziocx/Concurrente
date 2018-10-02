package main

import (
	"fmt"
	"math"
)

var end chan int

func max(s []float64) float64 {
	const h = 4
	b := int(math.Ceil(float64(len(s)) / float64(h)))

	m := make([]float64, h)

	for i := 0; i < h; i++ {
		ini := i * b
		fin := (i + 1) * b
		go func(idx, ini, fin int) {
			m[idx] = s[idx]
			for j := ini; j < fin && j < len(s); j++ {
				if s[j] > m[idx] {
					m[idx] = s[j]
				}
			}
			end <- 5
		}(i, ini, fin)
	}

	for i := 0; i < h; i++ {
		<-end
	}

	maxi := m[0]
	for _, v := range m {
		if v > maxi {
			maxi = v
		}
	}

	return maxi
}

func main() {
	end = make(chan int, 1)
	a := []float64{1, 6, 43, 6, 8, 43, 6, 35, 7, 94}
	fmt.Println(max(a))
}
