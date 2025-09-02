package main

import (
	"fmt"
)

func main() {
	var i int32 = 1
	for i <= 10 {
		fmt.Println("Table de ", i, " : ", i, "x1=", i*1, " ", i, "x2=", i*2, " ", i, "x3=", i*3, " ", i, "x4=", i*4, " ", i, "x5=", i*5)
		i = i + 1
	}
}
