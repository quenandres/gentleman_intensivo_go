package domain

import (
	"fmt"
	"learning-go-mustache/domain/entities"
)

func App() {
	persona := entities.Persona{
		Nombre:   "Juan",
		Apellido: "Perez",
		Edad:     20,
	}
	fmt.Println(persona.Saludar())
	persona.CumplirAnios()
	fmt.Println(persona.Saludar())

	elemento := entities.TipoElemento{
		Nombre: "Agua",
	}
	fmt.Println(elemento.Nombre)
}
