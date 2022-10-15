package main

import "fmt"

func main() {
	//Operadores Aritimetricos: + - / * %
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//Operadores 1Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores Relacionais
	fmt.Println(1 > 2)  //Maior
	fmt.Println(1 >= 2) //Maior ou igual
	fmt.Println(1 == 2) //Igual
	fmt.Println(1 <= 2) //Menor ou igual
	fmt.Println(1 < 2)  //Menor
	fmt.Println(1 != 2) //Diferente

	//Operadores Lógicos
	println("-----------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(verdadeiro && !false)

	//Operadores Unários
	n1 := 10
	n1++
	n1 += 10
	fmt.Println(n1)

	n2 := 10
	n2 *= 10
	n2 /= 5
	fmt.Println(n2)
	n2 %= 3
	fmt.Println(n2)

	//Operador Ternario não existe em golang
	//Em vez de utilizar texto := numero >5 ?"Maior que 5":"Menor que 5"
	//Utilizar
	var texto string
	if n2 > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
