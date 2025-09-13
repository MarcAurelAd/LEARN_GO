package main

//cet algorithme fusionne deux tableaux entrés par l'utilisateur et range les valeurs en ordre croissant

import (
	"fmt"
)

func main() {
	taille1 := 0
	taille2 := 0

	fmt.Print("Bonjour ! Veuillez entrer la taille des tableaux (Ex : X; X) : ")
	fmt.Scan(&taille1)
	fmt.Scanln(&taille2)

	tab1 := make([]int, taille1)
	tab2 := make([]int, taille2)

	fmt.Printf("Veuillez entrer les valeurs du tableau 1 de taille %d une à une (Ex: x; x; x;...): \n", taille1)
	for i := 0; i < taille1; i++ {
		fmt.Scan(&tab1[i])
	}

	fmt.Printf("Veuillez entrer les valeurs du tableau 2 de taille %d une à une (Ex: x; x; x;...): \n", taille2)
	for i := 0; i < taille2; i++ {
		fmt.Scan(&tab2[i])
	}
	//fusion des tableaux
	tab3 := append(tab1, tab2...)

	//tri du nouveau tableau
	taille3 := taille1 + taille2
	temp := 0
	i_temp := 0
	for i := 0; i < taille3; i++ {
		temp = tab3[i]
		i_temp = i
		for j := i; j < taille3; j++ {
			if tab3[j] < temp {
				temp = tab3[j]
				i_temp = j
			}
		}
		tab3[i_temp] = tab3[i]
		tab3[i] = temp
	}

	fmt.Printf("La fusion de vos deux tableaux donne le tableau suivant rangé en ordre croissant : %d \n", tab3)

}
