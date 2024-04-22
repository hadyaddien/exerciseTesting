package main_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volume", func() {
	var (
		panjangBalok   float64
		lebarBalok     float64
		tinggiBalok    float64
		jariJariTabung float64
		tinggiTabung   float64
	)

	BeforeEach(func() {
		panjangBalok = 5.0
		lebarBalok = 3.0
		tinggiBalok = 2.0
		jariJariTabung = 3.0
		tinggiTabung = 5.0
	})

	Describe("HitungVolumeBalok", func() {
		It("should calculate the volume of a box", func() {
			volume := HitungVolumeBalok(panjangBalok, lebarBalok, tinggiBalok)
			Expect(volume).To(Equal(30.0))
		})

		It("should calculate the volume of another box", func() {
			volume := HitungVolumeBalok(7.0, 4.0, 3.0)
			Expect(volume).To(Equal(84.0))
		})
	})

	Describe("HitungVolumeTabung", func() {
		It("should calculate the volume of a cylinder", func() {
			volume := HitungVolumeTabung(jariJariTabung, tinggiTabung)
			Expect(volume).To(BeNumerically("~", 141.37, 0.01))
		})

		It("should calculate the volume of another cylinder", func() {
			volume := HitungVolumeTabung(4.0, 6.0)
			Expect(volume).To(BeNumerically("~", 301.59, 0.01))
		})
	})
})

// Fungsi untuk menghitung volume balok
func HitungVolumeBalok(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

// Fungsi untuk menghitung volume tabung
func HitungVolumeTabung(jariJari, tinggi float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) * tinggi
}
