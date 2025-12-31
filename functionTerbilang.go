package main

import (
	"fmt"
	"strings"
)

func terbilang(n int) string {
	if n == 0 {
		return ""
	}

	angka := []string{
		"", "Satu", "Dua", "Tiga", "Empat",
		"Lima", "Enam", "Tujuh", "Delapan",
		"Sembilan", "Sepuluh", "Sebelas",
	}

	switch {
	case n < 12:
		return angka[n]

	case n < 20:
		return terbilang(n-10) + " Belas"

	case n < 100:
		return terbilang(n/10) + " Puluh " + terbilang(n%10)

	case n < 200:
		return "Seratus " + terbilang(n-100)

	case n < 1000:
		return terbilang(n/100) + " Ratus " + terbilang(n%100)

	case n < 2000:
		return "Seribu " + terbilang(n-1000)

	case n < 1000000:
		return terbilang(n/1000) + " Ribu " + terbilang(n%1000)

	case n < 1000000000:
		return terbilang(n/1000000) + " Juta " + terbilang(n%1000000)
	}

	return ""
}

func main() {
	x := 1152200
	result := strings.TrimSpace(terbilang(x) + "Rupiah")
	fmt.Println(result)
}
