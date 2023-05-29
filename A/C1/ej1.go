// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	getDiscount(200, 50)
}

func getDiscount(price int, perc int) {
	price -= (price * perc) / 100
	fmt.Println("El precio final es de: $" + strconv.Itoa(price) + ".")
}
