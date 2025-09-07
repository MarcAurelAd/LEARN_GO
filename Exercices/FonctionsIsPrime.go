package main

import "fmt"

// IsPrime vérifie si un nombre est premier.
func IsPrime(n int) bool {
	// Les nombres inférieurs ou égaux à 1 ne sont pas premiers.
	if n <= 1 {
		return false
	}
	// 2 est le seul nombre premier pair.
	if n == 2 {
		return true
	}
	// Les nombres pairs supérieurs à 2 ne sont pas premiers.
	if n%2 == 0 {
		return false
	}
	// On vérifie les diviseurs impairs jusqu'à la racine carrée de n.
	// Si on trouve un diviseur, le nombre n'est pas premier.
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	// Si aucun diviseur n'a été trouvé, le nombre est premier.
	return true
}

func main() {
	num1 := 0
	fmt.Print("Veuillez entrer un nombre : ")
	fmt.Scanln(&num1)
	fmt.Printf("%d est premier : %t\n", num1, IsPrime(num1))

}
