package main

import "fmt"

func main() {
	funcion()

	saludar("Michael")

	fmt.Println(sumar(7, 9))
	sumar2 := sumar(7, 9)
	fmt.Println(sumar2)
}

func funcion() {
	fmt.Println("Mi funcion")
}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}

func sumar(nro1 int, nro2 int) int {
	suma := nro1 + nro2

	return suma
}
