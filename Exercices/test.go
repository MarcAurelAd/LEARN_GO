package main

import "fmt"

func main() {
	chaineAvecEmojis := "Bonjour 😊"
	compteurRunes := 0
	for range chaineAvecEmojis {
		compteurRunes++
	}

	fmt.Printf("La chaîne \"%s\" contient %d caractères (runes) [9].\n", chaineAvecEmojis, compteurRunes)
}
