package main

import (
	"fmt"
)

func main() {
	const taille = 6
	tab1 := [taille]int64{0, 0, 0, 0, 0, 0}
	tab2 := [taille]int64{0, 0, 0, 0, 0, 0}

	fmt.Println("Veuillez entrer les valeurs de votre tableau : ")
	for i := 0; i < taille; i++ {
		fmt.Scanln(&tab1[i])
	}

	var temp int64
	var max int64 = tab1[0]
	var i_temp int
	//trouver la valeur max du tab1

	for i := 0; i < taille; i++ {

		if tab1[i] > max {
			max = tab1[i]
		}

	}

	//compare chaque valeur du tableau 1 à toutes les autres pour déterminer la plus petite
	for j := 0; j < taille; j++ {
		temp = tab1[j]
		i_temp = j
		for i := 0; i < taille; i++ {
			if tab1[i] < temp {
				temp = tab1[i]
				i_temp = i
			}

		}
		//remplace la plus petite valeur trouvée dans le tableau par la valeur max pour réduire
		tab2[j] = temp
		tab1[i_temp] = max
	}
	fmt.Print("Les valeurs de votre tableau sont en ordre croissant : ")
	for i := 0; i < taille; i++ {
		fmt.Print(tab2[i])
		if i < (taille - 1) {
			fmt.Print(", ")
		} else {
			fmt.Print(".")
		}
	}
}
