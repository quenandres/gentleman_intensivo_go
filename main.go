package main

import (
	"errors"
	"fmt"
)

func suma(a, b int) int {
	return a + b
}

//funcion que devuelve mas de 1 valor

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		error := errors.New("no se puede dividir por 0")
		return 0, error
	}

	conciente := a / b

	return conciente, nil
}

func main() {
	conciente, error := dividir(10, 0)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println("El conciente es: ", conciente)
}
