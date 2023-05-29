// Ejercicio 2 - Préstamo
// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar:
// solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
// se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
// Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.

package main

import "fmt"

func main() {
	evaluateClient(21, true, 19, 150000)

	evaluateClient(22, false, 0, 1000)

	evaluateClient(23, true, 20, 200000)
}

func evaluateClient(age int, employed bool, seniority int, salary float32) {
	approved := false
	if age > 22 {
		fmt.Println("Edad aprobada: mayor a 22 años.")
		approved = true
	} else {
		fmt.Println("Edad desaprobada: no es mayor a 22 años.")
		approved = false
	}
	if employed {
		fmt.Println("Empleo aprobado: tiene empleo.")
		approved = true

	} else {
		fmt.Println("Empleo desaprobada: está desempleado.")
		approved = false

	}
	if seniority > 12 {
		fmt.Println("Antiguedad aprobada: es mayor a 1 año.")
		approved = true

	} else {
		fmt.Println("Antiguedad desaprobada: no es mayor a 1 año.")
		approved = false
	}
	if approved {
		if salary > 100000 {
			fmt.Println("EXTRA Sueldo aprobado: es mayor a $100.000, no se cobrarían intereses.")
		} else {
			fmt.Println("EXTRA Sueldo desaprobado: no es mayor a $100.000, se cobrarían intereses.")
		}
	}

	fmt.Println("--------------------------------------------------------------------------")

}
