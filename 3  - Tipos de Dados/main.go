package main

import (
	"errors"
	"fmt"
)

func main() {
	//números  inteiros
	//Os tipos inteiros são int8,int16,int32,int64
	// usint é um tipo inteiro sem sinal. assim podendo ser apenas positivo
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// alias
	//int32 = rune

	var numero3 rune = 123456
	fmt.Println(numero3)

	//int8 = byte
	var numero4 byte = 127
	fmt.Println(numero4)

	// Números reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12845.9448934
	fmt.Println(numeroReal3)

	// Texto
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Caso declare uma variavel e defina a ela um caracteres com aspas únicas essa váriavel recebera o número correspondente ao caractere na tabela ascii
	x := 'B'
	fmt.Println(x)

	//Valor zero
	//String recebe ''
	//Int e float recebe 0
	//boolean recebe false
	//Error recebe nil

	//Boolean
	var boolean1 bool
	fmt.Println(boolean1)

	//Error
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
