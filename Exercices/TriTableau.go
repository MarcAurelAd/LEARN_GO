package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var err error
	var tab1 []int
	//var tab2 []int
	var max int = 0

	fmt.Println("Veuillez entrer les valeurs de votre tableau et entrer k pour achever la lecture du tableau : ")
	input := "dir"

	for {

		fmt.Scanln(&input)
		num, err = strconv.Atoi(input) // Conversion de la chaîne en entier

		// Vérification de l'erreur

		if err != nil {

			break
		}
		tab1 = append(tab1, num)

		//trouver la valeur max du tab1
		if num > max {
			max = num
		}

	}
	//test pour vérifier si le bloc de code pour gérer les entrées utilisateur fonctionne bien
	/*fmt.Println("Le maximum est", max)
	fmt.Println("Valeurs entrées : ", tab1)
	//var temp int
	//var i_temp int

	//compare chaque valeur du tableau 1 à toutes les autres pour déterminer la plus petite
	/*for j := 0; j < taille; j++ {
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
	}*/
}
