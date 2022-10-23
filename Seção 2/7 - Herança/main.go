package main

import "fmt"

/*
	Não existe herança em Go, o que se aproxima mais disso é passar um string dentro do
	outro, porém sem declarar o tipo, Como demonstrado abaixo.
*/
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança só que não")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome, "Tem", e1.idade, "anos e cursa", e1.curso)

}
