// Registro de estudiantes
// Una universidad necesita registrar a los estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]
// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos. Para ello es necesario generar una estructura Alumno con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle.

package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	a1 := Alumno{"Zoe", "Jimenez", 22098123, "28-05-2023"}
	a1.Detalle()
}

func (a *Alumno) Detalle() {
	fmt.Println("Nombre: ", a.Nombre)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
}
