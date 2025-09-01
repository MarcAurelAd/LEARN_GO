package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a string  = "bon"
		b string  = "bonne"
		c float64 = 0
		d float64 = 0
	)
	fmt.Print("Bonjour ! Veuillez entrer vos deux nombres un par ligne : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	//conversion des variables string "a" et "b" vers des variables float64 "c" et "d" respectivement

	c, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Erreur. Veuillez entrer un nombre.")
		return
	}

	d, errr := strconv.ParseFloat(b, 64)
	if errr != nil {
		fmt.Println("Erreur. Veuillez entrer un nombre.")
		return
	}

	if b == "0" {
		b = a
		a = "0"
	}
	var choice string = "Bon"
	fmt.Print("Veuillez entrer l'opération à effectuer parmi : \n", "1) Addition\n", "2) Soustraction\n", "3) Multiplication\n", "4) Division\n")
	fmt.Scanln(&choice)
	fmt.Print("Le résultat est : ")

	if choice == "1" {
		fmt.Print(c + d)
	}

	if choice == "2" {
		fmt.Print(c - d)
	}

	if choice == "3" {
		fmt.Print(c * d)
	}

	if choice == "4" {
		fmt.Print(c / d)
	}
}
