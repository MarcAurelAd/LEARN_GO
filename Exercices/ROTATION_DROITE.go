package main

//Ce programme déplace de deux places vers la droite les valeurs d'un tableau entré par l'utilisateur

import (
	"fmt"
)

func main() {
	fmt.Print("Bonjour ! Veuillez entrer la taille de votre tableau : ")
	taille := 0
	fmt.Scanln(&taille)
	fmt.Println("Veuillez entrer les valeurs de votre tableau une à une :")
	//make a slice with the numbers inputed by the user
	tab1 := make([]int, taille)
	for i := 0; i < taille; i++ {
		fmt.Scanln(&tab1[i])
	}
	fmt.Print("Veuillez entrer de combien de postions les valeurs seront déplacées : ")
	k := 0
	fmt.Scanln(&k)

	tab2 := make([]int, taille)
	new_index := 0
	for i := 0; i < taille; i++ {
		new_index = (i + k%taille)
		if new_index >= taille {
			for new_index >= taille {
				new_index = (new_index - taille)
			}
		}
		tab2[new_index] = tab1[i]
	}

	fmt.Printf("Votre tableau devient après rotation : %d", tab2)
}
