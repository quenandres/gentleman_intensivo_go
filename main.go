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

// funcion que recibe un numero indeterminado de argumentos
func imprimirNombres(nombres ...string) {

	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

// ejemplo de closure
func contador() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

type Rectangulo struct {
	Base, Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Altura * r.Base
}

func (r Rectangulo) Escalar(factor float64) Rectangulo {
	return Rectangulo{
		Base:   r.Base * factor,
		Altura: r.Altura * factor,
	}
}

func main() {
	conciente, error := dividir(10, 5)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println("El conciente es: ", conciente)

	imprimirNombres("Juan", "Pedro", "Maria")

	cont := contador()
	fmt.Println("Contador:", cont())
	fmt.Println("Contador:", cont())
	fmt.Println("Contador:", cont())

	rect := Rectangulo{Base: 10, Altura: 5}
	fmt.Println("Area:", rect.Area())

	escalado := rect.Escalar(3)
	fmt.Println("Base escalada:", escalado.Base, "Altura escalada:", escalado.Altura)
	fmt.Println("Area escalada:", escalado.Area())
}
