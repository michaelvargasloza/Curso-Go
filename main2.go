package main

import "fmt"

func main() {
	suma := 10 + 5
	resta := 10 - 5
	multi := 10 * 5
	div := 10 / 5
	resto := 10 % 3

	fmt.Print(suma, resta, multi, div, resto)

	suma++
	fmt.Println(suma)
	suma--
	fmt.Println(suma)
	suma += 10
	fmt.Println(suma)

	//Operadores tipo lógico
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 == 10)
	fmt.Println(10 != 5)

	fmt.Println(10 < 5 && 10 > 4)
	fmt.Println(10 > 5 && 10 > 4)

	fmt.Println(10 < 5 || 10 > 4)
}
