package main

import "fmt"

func Soma(numeros ...int) int {
	resultado := 0

	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

func main() {
	resultadoSoma := Soma(1, 52, 45, 18, 54, 45, 8, 7, 9, 4, 43, 7431, 163)
	fmt.Println(resultadoSoma)
}
