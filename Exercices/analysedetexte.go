package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bonjour ! Veuillez entrer votre phrase : ")
	phrase := "Hello"

	//syntax to use bufio scanner that also takes whitespaces as input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		phrase = scanner.Text()

	}

	longueur := len(phrase)

	//convertit la phrase en une suite de caractères, un caractère par mot dans la phrase initiale
	mots := strings.Fields(phrase)

	// Supprimer les espaces blancs autour de la chaîne de caractères (y compris le retour à la ligne)
	phrase = strings.TrimSpace(phrase)

	voyelles := "aeéèêïëiouyAEIOUY"
	nombreVoyelles := 0

	// Parcourir chaque caractère de l'input
	for _, char := range phrase {
		// Vérifier si le caractère est présent dans la chaîne des voyelles
		if strings.ContainsRune(voyelles, char) {
			nombreVoyelles++
		}

	}

	consonnes := "zrtpqsdghfjklmwxcvbnZRTPQSDFGHJKLMWXCVBN"
	nombreConsonnes := 0

	// Parcourir chaque caractère de l'input
	for _, char := range phrase {
		// Vérifier si le caractère est présent dans la chaîne des voyelles
		if strings.ContainsRune(consonnes, char) {
			nombreConsonnes++
		}

	}
	fmt.Println("Votre phrase contient ", longueur, "caractères, ", len(mots), "mots, ", nombreVoyelles, " voyelles et", nombreConsonnes, "consonnes.")

	//passe l'input en majuscule/miniscule
	phraseMaj := strings.ToUpper(phrase)
	phraseMin := strings.ToLower(phrase)

	fmt.Println("Vous avez entré\n", phraseMaj, "\n", phraseMin)

}
