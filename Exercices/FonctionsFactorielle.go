package main

import (
	"fmt"
)

// Factorielle calcule la factorielle d'un nombre entier non négatif.
// Elle renvoie 1 pour 0! et 1!
// Pour n > 1, elle renvoie n * Factorielle(n-1).
func Factorielle(n int) int {
	if n < 0 {
		// La factorielle n'est pas définie pour les nombres négatifs.
		// Dans un vrai programme, vous pourriez vouloir renvoyer une erreur.
		fmt.Println("Erreur : la factorielle n'est pas définie pour les nombres négatifs.")
		return -1 // Valeur indicative d'erreur
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorielle(n-1)
}

func main() {
	nombre := 0
	fmt.Print("Veuillez entrer un nombre : ")
	fmt.Scanln(&nombre)
	resultat := Factorielle(nombre)
	fmt.Printf("La factorielle de %d est %d\n", nombre, resultat)
}
