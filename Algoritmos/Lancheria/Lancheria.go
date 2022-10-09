package main

import "fmt"

func main() {

	var codigo, qtd float64

	fmt.Println("***** Cardapio ***** ")
	fmt.Println("COD(100)Cachorro quente Preço:1,20")
	fmt.Println("COD(101)Bauru simples Preço:1,30")
	fmt.Println("COD(102)Bauru com ovo Preço:1,50")
	fmt.Println("COD(103)Hambúrger Preço:1,20")
	fmt.Println("COD(104)Cheeseburguer Preço:1,30")
	fmt.Println("COD(105)Refrigerante Preço:1,00")

	fmt.Println("\nEntre com seu pedido:")
	fmt.Println("Codigo do produto:")
	fmt.Scan(&codigo)
	fmt.Println("Quantidade:")
	fmt.Scan(&qtd)

	switch codigo {
	case 100:
		fmt.Println("O valor total do lanche foi:", qtd*1.20, "reais")
	case 101:
		fmt.Println("O valor total do lanche foi:", qtd*1.30, "reais")
	case 102:
		fmt.Println("O valor total do lanche foi:", qtd*1.50, "reais")
	case 103:
		fmt.Println("O valor total do lanche foi:", qtd*1.20, "reais")
	case 104:
		fmt.Println("O valor total do lanche foi:", qtd*1.30, "reais")
	case 105:
		fmt.Println("O valor total do lanche foi:", qtd*1.00, "reais")
	default:
		fmt.Println("Atenção !! codigo invalido")
	}

}
