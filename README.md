# Go

Resuelve concurrencia y escalabilidad
Facil de aprender y de escribir
Altamente eficiente
Concurrencia incorporada
Nacio en el 2007, se populariza por la sencilles y eficiencia


### Usos
Servidores web
herramientas de linea (CLI)
Aplicaciones web
Servicios en la nube

# Paquetes

Se divide por paquetes

# Variables

bool true/false | flag o condicionales | false
string cadena de caracteres | representa textos | ""
int, int8, int16, int32, int64 | entero | controlar el tamaño de los enteros | 0
uint, uint8, uint16, uint32, uint64 | entero sin signo | cuando no queremos negativos | 0
float32, float64 | representar valores numericos reales | 32 o 64 según el sistema operativo | 0
byte | es un uint8 | trabajar con datos binarios | 0
rune int32 | cuando queremos representar 1 solo caracter que ocupa mas de un byte | 0
complex64, complex128 Cuando tiene una parte real y otra parte imaginaria | N + i | 0 + 0i




### Variables
```go
//Numeros
	entero := 10
	decimal := 10.5
	suma := entero + int(decimal)
	fmt.Println(suma)

	// Cadenas de caracteres
	mensaje := "Hola"
	concatenado := mensaje + " Mundo"
	enMayusculas := strings.ToUpper(concatenado)
	fmt.Println(enMayusculas)

	// Bool
	esVerdadero := true
	fmt.Println(esVerdadero)

	// Array y slice
	arrayFijo := [7]int{1, 4, 5, 6, 7, 8, 9}
	slice := []int{4, 5, 6, 7, 8, 9, 7}
	slice = append(slice, 10)

	fmt.Println(arrayFijo)
	fmt.Println(slice)

	// Map
	diccionario := map[string]int{
		"clave1": 1,
		"clave2": 2,
		"clave3": 3,
	}
	fmt.Println(diccionario)

	// Struct
	type Persona struct {
		Nombre string
		Edad   int
		Email  string
	}

	persona := Persona{
		Nombre: "Juan",
		Edad:   31,
		Email:  "juan@gmail.com",
	}
```



# Condicionales bucles
```go
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
```


# Funciones
```go
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

```

# Punteros
```go
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
```
# Interfaces
```go
package main

import (
	"fmt"
	"math"
)

// Las interfaces no se implementa, solo se cumplen

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f Forma) {
	fmt.Printf("El area es: %.2f\n", f.Area())
}

func main() {
	c := Circulo{Radio: 2}
	imprimirArea(c)
}

```

