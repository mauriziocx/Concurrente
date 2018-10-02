package main

import (
    "fmt"
)

const TH = 4
var ch chan bool

func matMul(a, b [][]float64) [][]float64 {
    na, ma := len(a), len(a[0])
    nb, mb := len(b), len(b[0])
    if (ma != nb) {
        return nil
    }
    c := make([][]float64, na)
    for i := range c {
        c[i] = make([]float64, mb)
    }
    for i := 0; i < na; i++ {
        for j := 0; j < mb; j++ {
            go func(i, j int) {
                sum := 0.
                for k := 0; k < ma; k++ {
                    sum += a[i][k] * b[k][j]
                }
                c[i][j] = sum
                ch <- true
            }(i, j)
        }
    }
    for i := 0; i < na; i++ {
        for j := 0; j < mb; j++ {
            <-ch
        }
    }
    return c
}

func matMul2(a, b [][]float64) [][]float64 {
    na, ma := len(a), len(a[0])
    nb, mb := len(b), len(b[0])
    if (ma != nb) {
        return nil
    }
    c := make([][]float64, na)
    for i := range c {
        c[i] = make([]float64, mb)
    }
    for th := 0; th < TH; th++ {
        go func(si0 int){
            for si := si0; si < na*mb; si += TH {
                i, j := si / mb, si % mb
                sum := 0.
                for k := 0; k < ma; k++ {
                    sum += a[i][k] * b[k][j]
                }
                c[i][j] = sum
            }
            ch <- true
        }(th)
    }
    for th := 0; th < TH; th++ {
        <-ch
    }
    return c
}

func main() {
    a := [][]float64 {{4, 3, 1, 2},
                      {5, 1, 3, 4},
                      {2, 2, 1, 3}}
    b := [][]float64 {{5, 5, 5},
                      {2, 1, 3},
                      {4, 2, 2},
                      {3, 2, 2}}
    ch = make(chan bool)
    c := matMul(a, b)
    fmt.Println(c)
    c2 := matMul2(a, b)
    fmt.Println(c2)
}
