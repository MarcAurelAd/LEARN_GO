package main

import (
	"fmt"
)

func main() {
	fmt.Println("+------------------+----------+----------+---------------+------------+")
	fmt.Println("|        NOM       |   MIN    |   MAX    |     MOYENNE   |   MEDIANE  |")
	fmt.Println("+------------------+----------+----------+---------------+------------+")

	fmt.Printf("|%-10s  |", "Richard BLAIREAU")
	fmt.Printf("%-10s|", "14")
	fmt.Printf("%-10s|", "17")
	fmt.Printf("%-10s     |", "15.5")
	fmt.Println("            |")

	fmt.Printf("|%-10s        |", "Emma STONE")
	fmt.Printf("%-10s|", "12")
	fmt.Printf("%-10s|", "20")
	fmt.Printf("%-10s     |", "16")
	fmt.Println("    16.25   |")

	fmt.Printf("|%-10s   |", "Alexandro LOPEZ")
	fmt.Printf("%-10s|", "15")
	fmt.Printf("%-10s|", "19")
	fmt.Printf("%-10s     |", "17")
	fmt.Println("            |")

	fmt.Println("+------------------+----------+----------+---------------+------------+")

}
