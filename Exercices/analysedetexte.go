package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("Votre phrase contient ", longueur, "caractères.")
}
