package main

import "fmt"

func main() {
	fmt.Println("Array Interno")

	//cria um array de 15 posições e retorna um slice referenciando as 10 primeiras posições
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) //capacidaden
}
