package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Gustavo"
	u.idade = 20
	fmt.Println(u)

	endereco1 := endereco{"Rua dos alfineiro", 4}

	usuario2 := usuario{"Gustavo", 21, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
