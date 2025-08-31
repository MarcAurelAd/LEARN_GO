package main

import (
	"fmt"
	"math/rand"
)

func main() {
	titre := "La mer rouge"
	auteur := "Alexandre"
	prix := rand.Int31n(30)
	disponibility := "en stock"
	fmt.Println("Les informations de ce livre sont : ")
	fmt.Println("Titre : ", titre)
	fmt.Println("Auteur : ", auteur)
	fmt.Println("Prix : ", prix, "$")
	fmt.Println("Disponibilité : ", disponibility)
	fmt.Println("Remise de 15% effectuée")
	fmt.Println("Prix final : ", prix-(prix*(15/100)), "$.")
}
