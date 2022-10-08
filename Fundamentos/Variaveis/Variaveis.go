package main

import "fmt"

// Declaração de variaveis em go
func main() {

	//Consegue declarar constantes dessa forma sendo primeiro o nome da const seguido do tipo e logo em seguida o valor
	//de atribuição
	const number float64 = 1.78
	const myName string = "Carolina Cedro"
	const myAge int = 27
	const myCountry string = "Brazil"

	fmt.Print("Meu nome é:", myName, "\n")
	fmt.Print("Minha idade é:", myAge, "\n")
	fmt.Print("Meu pais é:", myCountry, "\n")

	//Também uma forma muito utilizada hoje em dia é declarar o nome e logo em seguida o valor tornando bem mais
	//simplificado desta maneira:

	estado := "Goiás" //A linguagem faz a inferencia de tipo automaticamente
	phone := 87878787
	peso := 70.0
	/*
		Aqui segue um ponto de atenção pq em go é obrigatorio utilizar a variavel após a criação. diferente das outras
		linguagens não podemos criar e usar em um outro determinado momento tem que ser instantaneo caso contrario o
		programa irá apresentar erros
	*/
	fmt.Println("Eu moro no estado de:", estado, "meu telefone é:", phone, "e o meu peso é:", peso)

	//é possivél declarar varias variaveis na mesma linha exemplo:

	var ativo, desativo bool = true, false

	fmt.Println("está ativo:", ativo)
	fmt.Println("está ativo:", desativo)

	//outra maneira muito interessante de declarar é com varios tipos na mesma linha

	cor, portas, ano, preco := "amarelo", 4, 2022, 2789.0
	fmt.Println(cor, portas, ano, preco)

}
