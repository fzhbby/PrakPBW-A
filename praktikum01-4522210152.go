// anonymous function

// Menghitung luas dan keliling dengan anonymous function
package main

import (
	"fmt"
	"math"
)

func luas(r float64) float64 {
	return math.Pi * r * r
}

func keliling(r float64) float64 {
	return 2 * math.Pi * r
}

// Menampilkan hasil
func main() {
	jariJari := 5.0

	luas := luas(jariJari)
	keliling := keliling(jariJari)

	fmt.Println("Luas lingkaran:", luas)
	fmt.Println("Keliling lingkaran:", keliling)
}
