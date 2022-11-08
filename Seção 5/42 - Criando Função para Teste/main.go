package main

import (
	"fmt"
	"introducao-a-testes/enderecos"
)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(TipoEndereco)
}
