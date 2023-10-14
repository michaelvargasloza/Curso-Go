package main

import "fmt"

func main() {
	numero := 45

	if numero == 40 {
		fmt.Println("Ejecuto algo dentro del If")
	} else if numero == 41 {
		fmt.Println("Ejecuto algo dentro del Else If")
	} else {
		fmt.Println("Ejecuto algo dentro del Else")
	}

	edad := 30
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad < 18 && edad > 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("La edad no es válida")
	}
}
