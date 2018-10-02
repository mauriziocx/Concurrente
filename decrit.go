package main

import (
    "fmt"
    "time"
    "math/rand"
)

func pausa() {
    t := rand.Intn(200) + 50
    time.Sleep(time.Nanosecond*time.Duration(t))
}

var turn int = 1

func P() {
    for i := 0; i < 10; i++ {
            pausa()
        fmt.Printf("P NO Critica 1\n")
            pausa()
        fmt.Printf("P NO Critica 2\n")
            pausa()
        fmt.Printf("P NO Critica 3\n")
        for turn != 1 {
            pausa()
        }
            pausa()
        fmt.Printf("P Critica 1\n")
            pausa()
        fmt.Printf("P Critica 2\n")
            pausa()
        fmt.Printf("P Critica 3\n")
        turn = 2
    }
    end <- true
}
func Q() {
    for i := 0; i < 10; i++ {
            pausa()
        fmt.Printf("\tQ NO Critica 1\n")
            pausa()
        fmt.Printf("\tQ NO Critica 2\n")
            pausa()
        fmt.Printf("\tQ NO Critica 3\n")
        for turn != 2 {
            pausa()
        }
            pausa()
        fmt.Printf("\tQ Critica 1\n")
            pausa()
        fmt.Printf("\tQ Critica 2\n")
            pausa()
        fmt.Printf("\tQ Critica 3\n")
        turn = 1
    }
    end <- true
}

var end chan bool

func main() {
    end = make(chan bool)
    go P()
    go Q()
    <- end
    <- end
}
