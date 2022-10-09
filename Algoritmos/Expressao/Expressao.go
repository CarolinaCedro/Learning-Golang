package main

import "fmt"

func main() {

	var a, b, c float64

	fmt.Println("Entre com 3 valores inteiros positivos:")
	fmt.Scan(&a, &b, &c)

	r := (a + b) * 2
	s := (b + c) * 2
	d := (r + s) / 2

	fmt.Println("O resultado da expressão é:", d)

}
