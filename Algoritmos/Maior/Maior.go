package main

import "fmt"

/*
Elaborar um algoritmo que lê 3 valores a,b,c e os escreve. A seguir, encontre o maior dos 3 valores e o escreva com a mensagem : "É o maior
*/

func main() {

	var a, b, c float32
	fmt.Println("Entre com 3 valores: ")
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		fmt.Println(a, "é maior")
	} else if b > a && b > c {
		fmt.Println(b, "é maior")
	} else if c > a && c > b {
		fmt.Println(c, "é maior")
	}
}
