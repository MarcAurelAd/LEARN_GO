package main

import (
	"fmt"
)

func main() {
	fmt.Print("Bonjour ! Veuillez entrer la taille de vos deux tableaux (EX: x; x) : ")
	taille1 := 0
	taille2 := 0
	fmt.Scan(&taille1)
	fmt.Scan(&taille2)

	fmt.Printf("Veuillez entrer les valeurs du tableau 1 de taille %d trié en ordre croissant (Ex: x; x;...) : \n", taille1)
	tab1, is_sorted := Entree_tab(taille1)

	/*if is_sorted == false {
		fmt.Print("Ce tableau n'est pas trié. Veuillez entrer T pour le trier : ")
		ans := "tree"
		fmt.Scanln(&ans)
	}*/

	if is_sorted == false {
		tab1 = Sort(tab1, taille1)
	}

	fmt.Printf("Veuillez entrer les valeurs du tableau 2 de taille %d trié en ordre croissant (Ex: x; x;...) : \n", taille2)
	tab2, is_sorted := Entree_tab(taille2)
	if is_sorted == false {
		tab2 = Sort(tab2, taille2)
	}

	//Fusion des tableaux triés 1 et 2
	j := 0
	k := 0
	tab3 := make([]int, (taille1 + taille2))
	for i := 0; i < (taille1 + taille2); i++ {
		if j < taille1 && k < taille2 {
			if tab1[j] <= tab2[k] {
				tab3[i] = tab1[j]
				j++
			} else if k < taille2 {
				tab3[i] = tab2[k]
				k++
			}
		} else if k >= taille2 {
			tab3[i] = tab1[j]
			j++
		} else if j >= taille1 {
			tab3[i] = tab2[k]
			k++
		}

	}
	fmt.Printf("Votre nouveau tableau est : %d", tab3)
}

func Entree_tab(taille int) ([]int, bool) {
	/*
			-Prend en paramètre la taille (taille) d'un tableau;
		  	-Permet à l'utilisateur d'entrer des nombres pour remplir le tableau;
		  	-Retourne le tableau (tab) ainsi créé et indique dans "is_sorted" si le tableau est trié ou pas avec la syntaxe (tab, is_sorted)
	*/

	tab := make([]int, taille)
	is_sorted := true
	for i := 0; i < taille; i++ {
		fmt.Scan(&tab[i])
		if i >= 1 && tab[i] < tab[(i-1)] && is_sorted == true {
			is_sorted = false
		}
	}
	return tab, is_sorted
}

func Sort(tab []int, taille int) []int {
	/*
		Prend en paramètres un tableau (tab) et sa taille (taille)
		Retourne le tableau trié (tab)
	*/
	temp := 0
	i_temp := 0
	for i := 0; i < taille; i++ {
		temp = tab[i]
		i_temp = i
		for j := i; j < taille; j++ {
			if tab[j] <= temp {
				temp = tab[j]
				i_temp = j
			}
		}
		tab[i_temp] = tab[i]
		tab[i] = temp
	}
	return tab
}
