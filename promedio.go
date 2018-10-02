package main

import (
	"fmt"
)

func promedio(s []int) {
	for i := 0; i < len(s); i++ {
		suma += s[i]
	}
	ch <- true
}

var suma int
var ch chan bool

func main() {
	s := []int{9, 9, 5, 9, 5, 9, 5, 5, 7, 7, 7, 7}
	f := 3
	suma = 0
	ch = make(chan bool, 3)
	for th := 0; th < f; th++ {
		izq := th * (f + 1)
		der := len(s)/f + izq
		go promedio(s[izq:der])
	}
	for th := 0; th < f; th++ {
		<-ch
	}
	fmt.Println(suma / len(s))
}
