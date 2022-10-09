package main

import "fmt"

//Faça um algoritmo que leia as 3 notas de um aluno e calcule a média final
//deste aluno. Considerar que a média é ponderada e que o peso das notas é:
//2,3 e 5, respectivamente.

func main() {

	var nota1, nota2, nota3 float64

	fmt.Println("Entre com as 3 notas:")
	fmt.Scan(&nota1, &nota2, &nota3)
	media := ((nota1 * 2) + (nota2 * 3) + (nota3 * 5)) / 10
	fmt.Printf("Sua media final foi: %.2f", media)
}
