package main

import (
	"fmt"
)

func main() {
	var Tempcelsisus int32 = 0
	var duree int32 = 0
	var masse int32 = 0
	var Tempfahrenheit int32 = 0
	fmt.Print("Bonjour ! Veuillez entrer la température à convertir en Fahrenheit : ")
	fmt.Scanln(&Tempcelsisus)
	Tempfahrenheit = (Tempcelsisus*(9/5) + 32)
	fmt.Println("Votre température en Fahrenheit est : ", Tempfahrenheit)
	fmt.Print("Veuillez entrer une duree en secondes : ")
	fmt.Scanln(&duree)
	hours := int32(duree / 3600)
	minutes := int32((duree - hours*3600) / 60)
	seconds := int32((duree - hours*3600 - minutes*60))
	fmt.Println("Votre duree vaut : ", hours, "h", minutes, "m", seconds, "sec.")
	fmt.Print("Veuillez entre votre poids en kg : ")
	fmt.Scanln(&masse)
	fmt.Print("Veuillez entre votre taille en mètres : ")
	var taille int32 = 0
	fmt.Scanln(&taille)
	fmt.Println("Votre IMC est : ", masse/(taille*taille), "kg/m²")
}
