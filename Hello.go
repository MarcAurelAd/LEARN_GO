package main

import (
	"fmt"
)

func main() {
	var Usernamed string = "Marc-Aurel"
	var Userage float64 = 17.56
	var Dee bool = true
	if Dee == true {
		Userage = 17
	} else {
		Userage = 18
	}

	fmt.Println("Bonjour. Je suis", Usernamed, "et j'apprend le GO. J'ai ", Userage, " ans.")
}
