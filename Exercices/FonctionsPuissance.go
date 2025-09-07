package main

import (
	"fmt"
	"math"
)

// Puissance calcule la base élevée à la puissance de l'exposant.
func Puissance(base float64, exposant float64) float64 {
	return math.Pow(base, exposant)
}

func main() {
	base := 0.0
	exposant := 0.0
	fmt.Print("Veuillez entrer la base à élever : ")
	fmt.Scanln(&base)
	fmt.Print("Veuillez entrer l'exposant de la base : ")
	fmt.Scanln(&exposant)

	resultat := Puissance(base, exposant)
	fmt.Printf("%.2f élevé à la puissance %.2f est %.2f\n", base, exposant, resultat)
}
