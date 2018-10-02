package main

import (
    "fmt"
    "time"
)

func maxS(s []float32)  float32 {
    m := s[0]
    for i := range s {
        if s[i] > m {
            m = s[i]
        }
    }
    return m
}
/*
func max(s []float32)  float32 { // INCORRECTO
    m := s[0]
    bloque := len(s) / 2
    for i := 0; i < 2; i++ {
        go func(p, q int) {
            for i := p; i < q && i < len(s); i++ {
                if s[i] > m {
                    m = s[i]
                }
            }
        }(i * bloque, (i + 1) * bloque)
    }
    time.Sleep(time.Millisecond)
    return m
}*/

func max(s []float32)  float32 {
    m := []float32 {0, 0}
    bloque := len(s) / 2
    for i := 0; i < 2; i++ {
        go func(idx, p, q int) {
            m[idx] = s[p]
            for i := p; i < q && i < len(s); i++ {
                if s[i] > m[idx] {
                    m[idx] = s[i]
                }
            }
        }(i, i * bloque, (i + 1) * bloque)
    }
    time.Sleep(time.Millisecond)
    if m[0] > m[1] {
        return m[0]
    } else {
        return m[1]
    }
}

func main() {
    s := []float32 { 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 7 }
    fmt.Println(max(s))
}