package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var err error
	var tab1 []int

	var max int = 0

	fmt.Println("Veuillez entrer les valeurs de votre tableau et entrer k pour achever la lecture du tableau : ")
	input := "dir"
	taille := 0
	for {

		fmt.Scanln(&input)
		num, err = strconv.Atoi(input) // Conversion de la chaîne en entier

		// Vérification de l'erreur

		if err != nil {

			break
		}
		tab1 = append(tab1, num) //ajoute num au tableau
		taille++
		//trouver la valeur max du tab1
		if num > max {
			max = num
		}

	}

	//test pour vérifier si le bloc de code pour gérer les entrées utilisateur fonctionne bien
	/*fmt.Println("Le maximum est", max)
	fmt.Println("Valeurs entrées : ", tab1)
	//var temp int
	//var i_temp int*/

	//compare chaque valeur du tableau 1 à toutes les autres pour déterminer la plus petite
	temp := 0
	i_temp := 0
	tab2 := make([]int, taille)
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
	fmt.Printf("Les valeurs de votre tableau sont en ordre croissant : %d", tab2)

}
