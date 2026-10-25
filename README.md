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