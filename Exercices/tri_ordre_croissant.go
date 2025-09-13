package main

// Ce programme trie en ordre croissant les valeurs d'un tableau entré par l'utilisateur

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
	//prend les valeurs du tableau dans l'ordre, les coompares à toutes les autres pour trouver la plus petite, met cette petite valeur à l'avant du tableau et répète l'opération avec les autres valeurs
	temp := 0
	i_temp := 0
	for i := 0; i < taille; i++ {
		temp = tab1[i]
		i_temp = i
		for j := i; j < taille; j++ {
			if tab1[j] <= temp {
				temp = tab1[j]
				i_temp = j
			}
		}
		tab1[i_temp] = tab1[i]
		tab1[i] = temp
	}
	fmt.Println("Les valeurs de votre tableau en ordre croissant sont : ", tab1)
}
