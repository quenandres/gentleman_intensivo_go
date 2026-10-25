package main

import "fmt"

func main() {
	defer fmt.Println("Termina")
	//Condicional
	// Assertiva
	edad := 20
	if edad < 18 {
		fmt.Println("Eres menor de edad")
		return
	}

	fmt.Println("Eres mayor de edad")

	// Bucle clasico
	for i := 0; i < 10; i++ {
		//fmt.Println("El valor de i es: ", i)
		fmt.Printf("El valor de i es: %d\n", i)
	}

	// Bucle while
	j := 0
	for j < 10 {
		fmt.Printf("El valor de j es: %d\n", j)
		j++
	}

	// Bucle infinito
	s := 0
	for {
		s++

		if s == 5 {
			continue
		}

		fmt.Printf("El valor de s es: %d\n", s)
		if s == 10 {
			break
		}
	}

	// Bucle range
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for indice, valor := range numeros {
		fmt.Printf("El valor de indice es: %d y el valor es: %d\n", indice, valor)
	}

	// Switch
	switch edad {
	case 18:
		fmt.Println("Eres mayor de edad")
	case 17:
		fmt.Println("Eres menor de edad")
	default:
		fmt.Println("No tienes edad")
	}
}
