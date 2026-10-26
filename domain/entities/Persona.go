package entities

import "strconv"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p Persona) Saludar() string {
	return "Hola " + p.Nombre + " " + p.Apellido + " de " + strconv.Itoa(p.Edad) + " años"
}

func (p *Persona) CumplirAnios() {
	p.Edad++
}
