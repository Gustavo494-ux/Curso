package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado ")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() { //
	fmt.Println(AlunoEstaAprovado(7, 8))
}
