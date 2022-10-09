package main

import "fmt"

func main() {

	var idade int

	fmt.Println("Entre com sua idade:")
	fmt.Scan(&idade)

	switch idade {
	case 5, 6, 7:
		fmt.Println("infantil A ")
	case 8, 9, 10:
		fmt.Println("infantil B")
	case 11, 12, 13:
		fmt.Println("juvenil A")
	case 14, 15, 17:
		fmt.Println("juvenil B")
	default:
		fmt.Println("adulto")
	}

}
