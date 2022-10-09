package main

import "fmt"

/*
Calcule a média aritmética das 3 notas de um aluno e mostre, além do valor da média,
uma mensagem de "Aprovado", caso a média seja igual ou superior a 6, ou a mensagem "reprovado", caso contrário.
*/

func main() {

	var nota1, nota2, nota3 float64

	fmt.Println("Entre com 3 notas:")
	fmt.Scan(&nota1, &nota2, &nota3)

	media := (nota1 + nota2 + nota3) / 3
	fmt.Printf("Sua media é: %.2f \n", media)
	if media >= 6.0 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado!")
	}
}
