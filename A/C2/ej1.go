// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.
// Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).

package main

import "fmt"

func main() {
	fmt.Println("Empleado 1: $200000")
	fmt.Println("Se le descontará del sueldo: $", calculateTax(200000))
	fmt.Println("Empleado 2: $1000")
	fmt.Println("Se le descontará del sueldo: $", calculateTax(1000))
	fmt.Println("Empleado 3: $80000")
	fmt.Println("Se le descontará del sueldo: $", calculateTax(80000))
}

func calculateTax(salary float64) float64 {
	var tax float64
	if salary > 50000 {
		tax = (salary * 17) / 100
	} else if salary > 150000 {
		tax = (salary * 27) / 100
	} else {
		tax = 0.0
	}
	return tax
}
