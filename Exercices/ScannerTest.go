package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Vous avez entré : %q\n", line)
	}

	input := "Ceci est une chaîne de caractères pour tester."

	// Divise la chaîne en mots en utilisant l'espace blanc comme séparateur
	mots := strings.Fields(input)

	// Compte le nombre de mots dans la tranche
	nombreDeMots := len(mots)

	fmt.Printf("La chaîne contient %d mots.\n", nombreDeMots) // Affichera : La chaîne contient 7 mots.

	fmt.Println("Entrez du texte :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Lire l'input jusqu'au retour à la ligne

	// Supprimer les espaces blancs autour de la chaîne de caractères (y compris le retour à la ligne)
	input = strings.TrimSpace(input)

	voyelles := "aeiouAEIOU"
	nombreVoyelles := 0

	// Parcourir chaque caractère de l'input
	for _, char := range input {
		// Vérifier si le caractère est présent dans la chaîne des voyelles
		if strings.ContainsRune(voyelles, char) {
			nombreVoyelles++
		}
	}

	fmt.Printf("Le nombre de voyelles dans votre texte est : %d\n", nombreVoyelles)
}
