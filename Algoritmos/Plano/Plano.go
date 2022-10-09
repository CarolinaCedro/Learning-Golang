package main

import (
	"fmt"
	"math"
)

/*
1. Construa um algoritmo que, tendo como dados de entrada dois pontos quaisquer no plano, P(x1,y1) e P(x2,y2),
escreva a distância entre eles.
*/

func main() {

	var x1, y1, x2, y2 float64

	fmt.Println("Entre com dois pontos de entradas x1/y1:")
	fmt.Println("x1: ")
	fmt.Scan(&x1)
	fmt.Println("y1: ")
	fmt.Scan(&y1)

	fmt.Println("Entre com dois pontos de entradas x2/y2:")
	fmt.Println("x2: ")
	fmt.Scan(&x2)
	fmt.Println("y2: ")
	fmt.Scan(&y2)
	distancia := ((x1 - x2) * 2) + ((y1 - y2) * 2)
	saida := math.Sqrt(distancia)

	fmt.Println(saida)
}
