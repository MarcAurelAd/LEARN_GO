package main

import "fmt"

// pgcd calcule le Plus Grand Commun Diviseur de deux entiers non négatifs
// en utilisant l'algorithme d'Euclide.
func pgcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b // L'astuce de l'algorithme d'Euclide
	}
	return a
}

func main() {
	num1 := 0
	num2 := 0
	fmt.Print("Bonjour! Veuillez entrer le plus grand de vos deux nombres : ")
	fmt.Scanln(&num1)
	fmt.Print("Veuillez entrer le second nombre : ")
	fmt.Scanln(&num2)

	resultat := pgcd(num1, num2)
	fmt.Printf("Le PGCD de %d et %d est : %d\n", num1, num2, resultat)

}
