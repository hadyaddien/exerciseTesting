package main

import (
	"fmt"
	"math"
)

// HitungVolumeBalok menghitung volume balok
func HitungVolumeBalok(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

// HitungVolumeTabung menghitung volume tabung
func HitungVolumeTabung(jariJari, tinggi float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) * tinggi
}

func main() {
	// Menghitung volume balok
	panjangBalok := 5.0
	lebarBalok := 3.0
	tinggiBalok := 2.0

	volumeBalok := HitungVolumeBalok(panjangBalok, lebarBalok, tinggiBalok)
	fmt.Printf("Volume balok dengan panjang %.2f, lebar %.2f, dan tinggi %.2f adalah %.2f\n", panjangBalok, lebarBalok, tinggiBalok, volumeBalok)

	// Menghitung volume tabung
	jariJariTabung := 3.0
	tinggiTabung := 5.0

	volumeTabung := HitungVolumeTabung(jariJariTabung, tinggiTabung)
	fmt.Printf("Volume tabung dengan jari-jari %.2f dan tinggi %.2f adalah %.2f\n", jariJariTabung, tinggiTabung, volumeTabung)
}
