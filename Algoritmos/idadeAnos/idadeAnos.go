package main

import "fmt"

//. Faça um algoritmo que leia a idade de uma pessoa expressa em anos, meses e dias e mostre-a expressa apenas em dias.

func main() {
	var anos, meses, dias int

	fmt.Println("Entre com sua idade expressa em anos, meses e dias: ")
	fmt.Println("Anos:")
	fmt.Scan(&anos)
	fmt.Println("Meses:")
	fmt.Scan(&meses)
	fmt.Println("Dias:")
	fmt.Scan(&dias)

	idade := (anos * 365) + (meses * 30) + dias

	fmt.Println("Sua idade expressa em dias é:", idade)
}
