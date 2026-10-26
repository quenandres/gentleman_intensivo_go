package main

import "fmt"

//var a = 1 // NO

// * -> puntero a un espacio de memoria
// & -> direccion de memoria de una variable
// Copia los argumentos por valor, no por referencia

// referencia de nombre "a" que apunta a un espacio de memoria que contiene el valor 1
func incrementar(numero *int) {
	// Nunca modificamos el argumento, sino el elemento al que apunta el puntero
	*numero++
}

// Queremos modificar el elemento original

func main() {
	valor := 10
	fmt.Println("Valor antes de incrementar:", valor)
	incrementar(&valor)
	fmt.Println("Valor del puntero:", &valor)
	fmt.Println("Valor despues de incrementar:", valor)

	// new()
	puntero := new(int)
	fmt.Println("ubicación de memoria inicial con new():", puntero)
	fmt.Println("valor inicial con new():", *puntero)

}
