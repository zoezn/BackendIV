// Ejercicio 2 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente map, debemos imprimir la edad de Benjamin.
//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado, también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del map.

package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var legalAgeEmployees int

	for _, v := range employees {
		if v > 21 {
			legalAgeEmployees++
		}
	}

	fmt.Println(employees["Benjamin"])
	fmt.Println("Cantidad de empleados mayores a 21 años: ", legalAgeEmployees)
	employees["Federico"] = 25
	fmt.Println("Empleado nuevo: ", employees["Federico"])
	fmt.Println("Empleado a eliminar: Pedro")
	delete(employees, "Pedro")
	fmt.Println(employees)
}
