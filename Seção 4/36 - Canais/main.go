package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começa a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto //Mandando um valor para dentro do canal
		//fmt.Println(texto)
		time.Sleep((time.Second))
	}

	close(canal) //está fechando o canal,  assim ele não recebe nem envia mais
}
