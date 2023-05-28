// Ejercicio 1 - Calcular salario
// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
// Categoría C: su salario es de $1.000 por hora.
// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

package main

import (
	"errors"
	"fmt"
)

func main() {
	s, err := (calculateSalary(60, "A"))
	// s, err := (calculateSalary(60, "Z"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(s)

}

func calculateSalary(min int, category string) (float64, error) {
	var salary float64
	var hours = min / 60
	var err error

	switch category {
	case "C":
		salary = float64(1000 * hours)
	case "B":
		salary = float64(1500*hours) * 1.2
	case "A":
		salary = float64(3000*hours) * 1.5
	default:
		err = errors.New("invalid category")
	}
	return salary, err
}
