package main

import (
	"fmt"
)

func main() {
	var a string = "waiter"
	fmt.Println("Veuillez entrer : ")
	fmt.Scanln(&a)
	fmt.Println(a)
}
