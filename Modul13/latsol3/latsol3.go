package main

import "fmt"

func main() {
    var target, donasi1, donasi2, donasi3, donasi4 int 
    var total int
    var jumlahDonatur int
    var selesai bool
    
    fmt.Scan(&target)
    
    jumlahDonatur = 0
    total = 0
    
    for selesai = false; !selesai; {
        jumlahDonatur++
        
        if jumlahDonatur == 1 {
            fmt.Scan(&donasi1)
            total += donasi1
        } else if jumlahDonatur == 2 {
            fmt.Scan(&donasi2)
            total += donasi2
        } else if jumlahDonatur == 3 {
            fmt.Scan(&donasi3)
            total += donasi3
        } else if jumlahDonatur == 4 {
            fmt.Scan(&donasi4)
            total += donasi4
        }
        
        selesai = total >= target
    }
    
    total = 0 
    
    if jumlahDonatur >= 1 {
        total += donasi1
        fmt.Printf("Donatur 1: Menyumbang %d. Total terkumpul: %d\n", donasi1, total)
    }
    if jumlahDonatur >= 2 {
        total += donasi2
        fmt.Printf("Donatur 2: Menyumbang %d. Total terkumpul: %d\n", donasi2, total)
    }
    if jumlahDonatur >= 3 {
        total += donasi3
        fmt.Printf("Donatur 3: Menyumbang %d. Total terkumpul: %d\n", donasi3, total)
    }
    if jumlahDonatur >= 4 {
        total += donasi4
        fmt.Printf("Donatur 4: Menyumbang %d. Total terkumpul: %d\n", donasi4, total)
    }
    
    fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", total, jumlahDonatur)
}
