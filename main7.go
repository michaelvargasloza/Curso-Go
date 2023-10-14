package main

import "fmt"

func main() {
	frutas := []string{"Manzana", "Banana", "Naranja"}

	fmt.Println(frutas[0])

	frutas = append(frutas, "Sandía", "Frutilla")

	frutas[0] = "Papaya"

	for i := 0; i < 5; i++ {
		fmt.Println(frutas[i])
	}

	for i := 0; i < len(frutas); i++ {
		if frutas[i] == "Frutilla" {
			fmt.Println("Hay una coincidencia")
		}
	}
}
