package main

import (
    "fmt"
    "math"
)

func main() {
    var bilangan float64
    var selesai bool
    
    fmt.Scan(&bilangan)
    
    saatIni := bilangan
    batasAtas := math.Ceil(bilangan)
    
    for selesai = false; !selesai; {
        fmt.Printf("%.1f\n", saatIni)
        saatIni += 0.1
        saatIni = math.Round(saatIni*10) / 10
        selesai = saatIni > batasAtas
    }
}
