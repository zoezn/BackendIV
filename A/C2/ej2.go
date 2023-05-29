// Ejercicio 2 - Calcular promedio
// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.

package main

import "fmt"

func main() {
	fmt.Println(calculateAverage(7, 7, 7, 7, 45, 2, 2, 4, 52, 1, 0, 0, 0, 0))

}

func calculateAverage(values ...int) int {
	var sum int
	var acc int
	for _, v := range values {
		if v > 0 {
			sum += v
			acc++
		}
	}
	return sum / acc
}
