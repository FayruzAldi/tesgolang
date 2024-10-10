package main

import (
	"fmt"
)

// calcPriceOfExcessBaggage menghitung biaya bagasi berlebih berdasarkan jenis penerbangan
func calcPriceOfExcessBaggage(userBaggage int, freeQuotaBaggage int, domesticPrice int, internationalPrice int, flightType string) int {
	excessBaggage := userBaggage - freeQuotaBaggage
	if excessBaggage <= 0 {
		return 0
	}

	if flightType == "domestic" {
		return excessBaggage * domesticPrice
	} else if flightType == "international" {
		return excessBaggage * internationalPrice
	}
	return 0
}

// calcFlightTime menghitung waktu penerbangan berdasarkan peta dunia dan negara asal serta tujuan
func calcFlightTime(worldMap []string, fromCountry string, destinationCountry string) int {
	fromIndex := -1
	toIndex := -1

	// Mencari index negara asal dan negara tujuan di dalam world map
	for i, country := range worldMap {
		if country == fromCountry {
			fromIndex = i
		}
		if country == destinationCountry {
			toIndex = i
		}
	}

	if fromIndex == -1 || toIndex == -1 {
		// Jika negara tidak ditemukan, kembalikan 0
		return 0
	}

	// Menghitung jarak antara dua negara
	distance := toIndex - fromIndex
	if distance < 0 {
		distance = -distance // Ambil nilai absolut dari jarak
	}

	return distance
}

func main() {
	// Test untuk calcPriceOfExcessBaggage
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "domestic"))      // Output: 1100000
	fmt.Println(calcPriceOfExcessBaggage(50, 40, 110000, 160000, "international")) // Output: 1600000
	fmt.Println(calcPriceOfExcessBaggage(40, 50, 110000, 160000, "international")) // Output: 0

	// Test untuk calcFlightTime
	worldMap1 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap1, "ID", "KH")) // Output: 4

	worldMap2 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap2, "SG", "PH")) // Output: 3

	worldMap3 := []string{"PH", "", "ID", "SG", "MY", "VN", "KH"}
	fmt.Println(calcFlightTime(worldMap3, "KH", "PH")) // Output: 6
}
