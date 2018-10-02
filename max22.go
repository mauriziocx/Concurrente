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
    
    xx := func(idx int, ss []float64) {
        m[idx] = ss[0]
        for _, v := range ss {
            if v > m[idx] {
                m[idx] = v
            }
        }
    }
    
    for i := 0; i < h; i++ {
        ini := i*b
        fin := (i+1)*b
        go xx(i, s[ini:fin])
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
    a := []float64{ 1, 6, 43, 6, 8, 43, 6 ,35, 7, 94 }
    fmt.Println(max(a))
}