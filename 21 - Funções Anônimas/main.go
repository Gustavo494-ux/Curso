package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido  -> %s", texto)
	}("Olá")

	fmt.Println(retorno)
}
