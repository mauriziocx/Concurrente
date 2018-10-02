package main

import (
    "fmt"
    "time"
    "math/rand"
)

func P() {
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(20) + 80))
    k1 := 1
    n = k1
}

func Q() {
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(20) + 80))
    k2 := 2
    n = k2
}

var n int

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    go P()
    go Q()
    time.Sleep(time.Second)
    fmt.Println(n)
}
