package main

import (
    "fmt"
)

func serialMatMul(a, b [][]float64) [][]float64 {
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
            sum := 0.
            for k := 0; k < ma; k++ {
                sum += a[i][k] * b[k][j]
            }
            c[i][j] = sum
        }
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
    c := serialMatMul(a, b)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}