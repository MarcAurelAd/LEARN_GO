package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a string = "bon"
	)
	fmt.Print("Bonjour ! Veuillez entrer vos deux nombres : ")
	fmt.Scanln(&a)
	valeur, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Erreur de conversion :", err)
		return
	}
	fmt.Println(valeur)
}
