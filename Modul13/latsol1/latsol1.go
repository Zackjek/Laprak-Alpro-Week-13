package main

import (
    "fmt"
)

func main() {
    var bilangan int
    
    fmt.Scan(&bilangan)

    jumlahDigit := 0

    for {
        jumlahDigit++
        bilangan /= 10
        if bilangan == 0 {
            break
        }
    }

    fmt.Print(jumlahDigit)
}
