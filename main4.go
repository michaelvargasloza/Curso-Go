package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Michael"

	fmt.Println(string(nombre[0]))

	for j := 0; j < 7; j++ {
		fmt.Println(j, string(nombre[j]))
	}

	//No existe la opción While
	numero := 1

	for numero != 10 {
		fmt.Println("Numero:", numero)
		numero++
	}
}
