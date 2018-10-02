package main

import (
    "fmt"
)

type Semaphore chan bool

func (s Semaphore) Wait() {
    s<- true
}
func (s Semaphore) Signal() {
    <-s
}

const N = 5
var fork []Semaphore
var room Semaphore
var end chan bool

func philosopher(id int) {
    izq := id
    der := (id+1) % N
    for i := 0; i < 10; i++ {
        fmt.Printf("Filósofo %d pensando\n", id)
        room.Wait()
        fork[izq].Wait()
        fork[der].Wait()
        fmt.Printf("Filósofo %d comiendo\n", id)
        fork[izq].Signal()
        fork[der].Signal()
        room.Signal()
    }
    end<- true
}

func main() {
    fork = make([]Semaphore, N)
    room = make(Semaphore, N-1)
    end = make(chan bool)
    for i := 0; i < N; i++ {
        fork[i] = make(Semaphore, 1)
    }
    for i := 0; i < N; i++ {
        go philosopher(i)
    }
    for i := 0; i < N; i++ {
        <-end
    }
}
