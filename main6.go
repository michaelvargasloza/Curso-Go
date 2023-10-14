package main

import "fmt"

// POO
type persona struct {
	nombre   string
	apellido string
	edad     int
}

// Métodos
func (p persona) saludar(saludo string) {
	fmt.Println(saludo + ", " + p.nombre)
}

func (pers persona) cumple() int {
	return pers.edad + 1
}

func main() {
	persona1 := persona{"Michael", "Vargas", 30}
	persona2 := persona{"Bill", "Gates", 65}
	fmt.Println("Persona1", persona1)
	fmt.Println("Persona2", persona2)

	persona1.edad = 31
	fmt.Println("Persona1", persona1)

	persona1.saludar("Hola")
	persona2.saludar("Hello")

	fmt.Println(persona1.cumple())

	edad_persona2 := persona2.cumple()
	fmt.Println(edad_persona2)
}
