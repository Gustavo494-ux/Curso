package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	variavel1 += 30
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referencia de memoria
	var variavel3 int
	var ponteio *int

	variavel3 = 100
	ponteio = &variavel3 //Referenciação
	// *ponteiro desrefenciação
	fmt.Println(variavel3, *ponteio)
}
