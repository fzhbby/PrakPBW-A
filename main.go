package main

import (
	"fmt"
	"progo/ratarata" //Menginpor package
)

func main() {
	//Data nilai siswa
	nilaiSiswa := []int{80, 75, 90, 85, 60}

	//Menghitung rata-rata dengan dungsi dari package
	rataRata := ratarata.HitungRataRata(nilaiSiswa)

	//Menampilkan hasil
	fmt.Println("Rata-rata nilai siswa:", rataRata)
}
