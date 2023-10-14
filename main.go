package main

import "fmt"

//Para ejecutar sin el IDE Goland (de paga) ejecutar los comandos:
//go build nombre_archivo.go
//./nombre archivo
//desde CMD ejecutar go version para verificar si esta instalado GO

func main() {
	var entero int
	var decimal float64
	var texto string
	var booleana bool

	entero = 5
	decimal = 10.4
	texto = "Hola Mundo"
	booleana = true

	fmt.Println(entero, decimal, texto, booleana)

	numero := 10
	fmt.Println(numero)
	//numero = "hola"
	//var entero string
}
