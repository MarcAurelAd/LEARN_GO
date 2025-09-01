package main

import (
	"fmt"
)

func main() {
	fmt.Println("+----------------+------+-----------+----------+")
	fmt.Println("|       NOM      |  AGE |  MOYENNE  |  STATUT  |")
	fmt.Println("+----------------+------+-----------+----------+")

	fmt.Printf("|%-10s    |", "Alice DUPONT")
	fmt.Print(" 20   |")
	fmt.Printf(" %10s|", "15.75")
	fmt.Printf("%10s|\n", "Admis")

	fmt.Printf("|%-10s      |", "Bob Martin")
	fmt.Print(" 19   |")
	fmt.Printf(" %10s|", "12.50")
	fmt.Printf("%10s|\n", "Refusé")

	fmt.Printf("|%-10s   |", "Claire MOREAU")
	fmt.Print(" 21   |")
	fmt.Printf(" %10s|", "17.25")
	fmt.Printf("%10s|\n", "Admis")

	fmt.Println("+----------------+------+-----------+-----------+")
\t
}
