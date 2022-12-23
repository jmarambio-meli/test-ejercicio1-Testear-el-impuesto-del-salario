package main

import "fmt"

func main() {
	var sueldo float64
	fmt.Println("Ingrese el sueldo porfavor")
	fmt.Scanln(&sueldo)
	impuesto := CalcularImpuesto(sueldo)
	fmt.Printf("Su impuesto es de: %v,\npor lo tanto el sueldo a recibir es de : %.2f\n", impuesto, sueldo-impuesto)
}

func CalcularImpuesto(sueldo float64) float64 {
	var impuesto float64

	switch {
	case sueldo > 50000 && sueldo < 150000:
		impuesto = sueldo * 0.17
	case sueldo > 150000:
		impuesto = sueldo * 0.27
	case sueldo < 50000:
		impuesto = 0
	}
	return impuesto
}
