// Calcular cantidad de alimento
// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hámsteres, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
// Tienen los siguientes datos:
// Perro: 10 kg de alimento.
// Gato: 5 kg de alimento.
// Hamster: 250 g de alimento.
// Tarántula: 150 g de alimento.
// Se solicita:
// Implementar una función animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso de que no exista el animal).
// Una función para cada animal que calcule la cantidad de alimento basándose en la cantidad del tipo de animal especificado.

package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	funcion, err := calculateFood("tarantula")
	if err != nil {
		log.Fatal(err)
	}
	cant := funcion(2)
	fmt.Println(cant)

}
func catCalc(cant int) int {
	return cant * 5000
}
func dogCalc(cant int) int {
	return cant * 10000
}
func hamsterCalc(cant int) int {
	return cant * 250
}
func tarantulaCalc(cant int) int {
	return cant * 150
}
func calculateFood(animal string) (func(int) int, error) {
	switch animal {
	case "dog":
		return dogCalc, nil
	case "cat":
		return catCalc, nil
	case "hamster":
		return hamsterCalc, nil
	case "tarantula":
		return tarantulaCalc, nil
	default:
		return nil, errors.New("invalid animal")
	}
}
