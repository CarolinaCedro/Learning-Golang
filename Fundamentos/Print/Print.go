package main

import "fmt"

func main() {

	//Formas de mostrar mensagens na tela

	nome := "Lalisa Manobal"
	cor := "Vermelho"
	ano := 2022
	frase := "Rapadura é doce mas num é mole né bem que o povo da igreja fala rs"
	peso := 57.90908
	ativo := true

	fmt.Println("Nome:", nome)
	fmt.Print("Nome:", nome)
	fmt.Print("\n")
	fmt.Printf("Nome: %s e peso: %.2f \n", nome, peso)
	fmt.Println(cor, ano, frase, ativo)
}
