// Ejercicio 1 - Listado de nombres
// Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista. Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
// Luego de dos clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente. El nombre de la nueva estudiante es Gabriela.

package main

import (
	"fmt"
)

var students []string

func main() {
	nombres := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	for i := 0; i < len(nombres); i++ {
		addStudents(nombres[i])
	}
	addStudents("Gabriela")
	fmt.Println(students)

}

func addStudents(name string) {
	students = append(students, name)
}
